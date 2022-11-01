<div align="center">
  <img alt="errdetails logo" src="assets/go.png" height="150" />

  <h3>Go Error Details</h3>
  <p>Enrich errors by context information.</p>
</div>

---

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/dnozdrin/errdetail)](/LICENSE)
[![Codecov](https://codecov.io/gh/dnozdrin/errdetail/branch/main/graph/badge.svg)](https://codecov.io/gh/dnozdrin/errdetail)
[![Coreportcard](https://goreportcard.com/badge/github.com/dnozdrin/errdetail)](https://goreportcard.com/report/github.com/dnozdrin/errdetail)
[![Go Reference](https://pkg.go.dev/badge/github.com/dnozdrin/errdetail.svg?style=flat-square)](https://pkg.go.dev/github.com/dnozdrin/errdetail)
[![Release](https://img.shields.io/github/release/dnozdrin/errdetail.svg)](https://github.com/dnozdrin/errdetail/releases/latest)

## Usage

### Create a detailed error

```go
err := errdetail.New(
    "bad request",
    errdetail.NewDetail(
        errdetail.WithDomain("user.auth"),
        errdetail.WithCode("invalid_email"),
        errdetail.WithDescription("email validation failed"),
        errdetail.WithField("user.email"),
        errdetail.WithReason("invalid character detected: \"#\""),
    )
)
```

### Wrap existing error

```go
err := errdetail.Wrap(
    errdetail.ErrInvalidArgument,
    "bad request",
    errdetail.NewDetail(
        errdetail.WithDomain("user.auth"),
        errdetail.WithCode("invalid_email"),
        errdetail.WithDescription("email validation failed"),
        errdetail.WithField("user.email"),
        errdetail.WithReason("invalid character detected: \"#\""),
    ),
    errdetail.NewDetail(
        errdetail.WithDomain("user.auth"),
        errdetail.WithCode("invalid_password"),
        errdetail.WithDescription("password validation failed"),
        errdetail.WithField("user.password"),
        errdetail.WithReason("password is empty"),
    ),
)
```

### Use provided error constructors

```go
err := NewNotFound("discount not found", errdetail.NewDetail(errdetail.WithCode("order_discount_not_supported")))
```

### Use predefined errors

```go
func (s *Service) GetCached(ctx context.Context, id uuid.UUID) (Item, error) {
    if item, ok := s.cache.Get(ctx, id); !ok {
        return Item{}, errdetail.ErrNotFound
    }
    
    // ...
}
```


### Transform errors details to a suitable presentation

```go

type ErrorResponse struct {
    Error *Error `json:"error,omitempty"`
}

type Error struct {
    Status int `json:"status"`
    Title string `json:"title"`
    Code ResponseCode `json:"code"`
    Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
    Domain string `json:"domain,omitempty"`
    Reason string `json:"reason,omitempty"`
    Field string `json:"field,omitempty"`
    Description string `json:"description,omitempty"`
    Code string `json:"code,omitempty"`
}

func NewErrorResponse(err error) ErrorResponse {
    if err == nil {
        return ErrorResponse{}
    }
    
    var (
        status int
        title  string
        code   ResponseCode
    )
    
    switch {
    case errors.Is(err, errdetail.ErrInvalidArgument):
        status = http.StatusBadRequest
        title = errdetail.ErrInvalidArgument.Error()
        code = "INVALID_ARGUMENT"
	
    case errors.Is(err, errdetail.ErrNotFound):
        status = http.StatusNotFound
        title = errdetail.ErrNotFound.Error()
        code = "NOT_FOUND"
    
	// ...
    
    default:
        status = http.StatusInternalServerError
        title = "unknown"
        code = "UNKNOWN"
    }
    
    return ErrorResponse{
        Error: &Error{
            Code:    code,
            Title:   title,
            Status:  status,
            Details: extractDetails(err),
        },
    }
}

func extractDetails(err error) []ErrorDetail {
    extracted := errdetail.ExtractDetails(err)
    if len(extracted) == 0 {
        return nil
    }
    
    details := make([]ErrorDetail, len(extracted))
        for i := range extracted {
            details[i] = ErrorDetail{
            Domain:      extracted[i].Domain(),
            Reason:      extracted[i].Reason(),
            Field:       extracted[i].Field(),
            Description: extracted[i].Description(),
            Code:        extracted[i].Code(),
        }
    }
    
    return details
}

```

`ErrorResponse` encoded to JSON:

```json
{
  "error": {
    "status": 400,
    "title": "invalid argument",
    "code": "INVALID_ARGUMENT",
    "details": [
      {
        "domain": "user.auth",
        "code": "invalid_email",
        "description": "email validation failed",
        "field": "user.email",
        "reason": "invalid character detected: \"#\""
      },
      {
        "domain": "user.auth",
        "code": "invalid_password",
        "description": "password validation failed",
        "field": "user.password",
        "reason": "password is empty"
      }
    ]
  }
}
```

For further details see [examples](https://github.com/dnozdrin/errdetail/tree/main/examples) and [reference](https://pkg.go.dev/badge/github.com/dnozdrin).

## Contributing

For contribution guidelines check [CONTRIBUTING.md](https://github.com/dnozdrin/errdetail/blob/main/CONTRIBUTING.md)
