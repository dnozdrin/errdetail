<div align="center">
  <img alt="errdetail logo" src="assets/go.png" height="150" />

  <h3>Errdetail</h3>
  <p>Enrich Go errors by context information.</p>
</div>

---

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/dnozdrin/errdetail)](/LICENSE)
[![Codecov](https://codecov.io/gh/dnozdrin/errdetail/branch/main/graph/badge.svg)](https://codecov.io/gh/dnozdrin/errdetail)
[![Coreportcard](https://goreportcard.com/badge/github.com/dnozdrin/errdetail)](https://goreportcard.com/report/github.com/dnozdrin/errdetail)
[![GitHub CI](https://github.com/dnozdrin/errdetail/actions/workflows/ci.yml/badge.svg)](https://github.com/dnozdrin/errdetail/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dnozdrin/errdetail.svg?style=flat-square)](https://pkg.go.dev/github.com/dnozdrin/errdetail)
[![Release](https://img.shields.io/github/release/dnozdrin/errdetail.svg)](https://github.com/dnozdrin/errdetail/releases/latest)

## Description

**Errdetail** allows to add one or multiple details to an `error`. Each detail may contain the next optional fields:

| Field       | Possible usage                                                                                                 | Example                                                           |
|-------------|----------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------|
| domain      | a particular interest, activity, or type of knowledge                                                          | `user.management`                                                 |
| code        | a unique code for a particular type of error                                                                   | `validation_email_invalid_character`                              |
| description | a short, human-readable summary of the problem that should change from occurrence to occurrence of the problem | `email address must follow the format described in the RFC 5322`  |
| field       | a field where the error occurred, common for validation issues                                                 | `user.email`                                                      |
| reason      | a human-readable explanation specific to this occurrence of the problem                                        | `an invalid character has been detected in the provided sequence` |
| meta        | arbitrary data required for the error explanation                                                              | `{"dummyField1":"dummyValue", "dummyField2":2}`                   |

With such approach it's possible to aggregate information about several errors into one type that implements the `error` interface.

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
        errdetail.WithMeta(errdetail.Meta{
			"dummyField1": "dummyValue",
			"dummyField2": 2,
        }),
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
        errdetail.WithMeta(errdetail.Meta{
            "dummyField1": "dummyValue",
            "dummyField2": 2,
        }),
    ),
    errdetail.NewDetail(
        errdetail.WithDomain("user.auth"),
        errdetail.WithCode("invalid_password"),
        errdetail.WithDescription("password validation failed"),
        errdetail.WithField("user.password"),
        errdetail.WithReason("password is empty"),
        errdetail.WithMeta(errdetail.Meta{
            "dummyField3": "hello world!",
            "dummyField4": 4,
        }),
    ),
)
```

### Use provided error constructors

```go
err := NewNotFound("discount not found", errdetail.NewDetail(errdetail.WithCode("order_discount_not_supported")))
```

### Use predefined errors

```go
func (a *Adapter) Get(ctx context.Context, id uuid.UUID) (Item, error) {
    if item, ok := a.storage.Get(ctx, id); !ok {
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
        "reason": "invalid character detected: \"#\"",
        "meta": {
          "link": "https://example.com",
          "translations": {
            "en": "Hello world!",
            "ua": "Привіт, світе!"
          }
        }
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
