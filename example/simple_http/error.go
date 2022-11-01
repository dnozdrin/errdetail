// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package simple_http

import (
	"context"
	"errors"
	"net/http"

	"github.com/dnozdrin/errdetail"
)

// ErrorResponse represents an HTTP server response in case of an error.
type ErrorResponse struct {
	Error *Error `json:"error,omitempty"`
}

// Error is expected to be filled only via NewErrorResponse.
// Intentionally omits wrapped error's message presentation.
type Error struct {
	// Status is the HTTP status code applicable to this problem.
	Status int `json:"status"`
	// Title is a short, human-readable summary of the problem that should not change
	// from occurrence to occurrence of the problem.
	Title string `json:"title"`
	// Code is an application-specific error code, expressed as a string value.
	Code ResponseCode `json:"code"`
	// Details represents explanations specific to this occurrence of the problem.
	Details []ErrorDetail `json:"details,omitempty"`
}

// ErrorDetail is expected to be filled only via NewErrorResponse.
type ErrorDetail struct {
	// Domain allows to determine the application features that case the error.
	Domain string `json:"domain,omitempty"`
	// Reason may provide explanations of the problem cause and possible ways to resolve it.
	Reason string `json:"reason,omitempty"`
	// Field represents a field name, where the problem occurred, commonly used
	// for validation errors.
	Field string `json:"field,omitempty"`
	// Description represents explanations specific to this part of the problem.
	Description string `json:"description,omitempty"`
	// Code is an application-specific code specific to this part of the problem,
	// expressed as a string value.
	Code string `json:"code,omitempty"`
}

// NewErrorResponse creates a ErrorResponse with properly filled fields.
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
		code = statusInvalidArgument

	case errors.Is(err, errdetail.ErrFailedPrecondition):
		status = http.StatusBadRequest
		title = errdetail.ErrFailedPrecondition.Error()
		code = statusFailedPrecondition

	case errors.Is(err, errdetail.ErrOutOfRange):
		status = http.StatusBadRequest
		title = errdetail.ErrOutOfRange.Error()
		code = statusOutOfRange

	case errors.Is(err, errdetail.ErrUnauthenticated):
		status = http.StatusUnauthorized
		title = errdetail.ErrUnauthenticated.Error()
		code = statusUnauthenticated

	case errors.Is(err, errdetail.ErrNotFound):
		status = http.StatusNotFound
		title = errdetail.ErrNotFound.Error()
		code = statusNotFound

	case errors.Is(err, errdetail.ErrAborted):
		status = http.StatusConflict
		title = errdetail.ErrAborted.Error()
		code = statusAborted

	case errors.Is(err, errdetail.ErrAlreadyExists):
		status = http.StatusConflict
		title = errdetail.ErrAlreadyExists.Error()
		code = statusAlreadyExists

	case errors.Is(err, errdetail.ErrRemoved):
		status = http.StatusGone
		title = errdetail.ErrRemoved.Error()
		code = statusRemoved

	case errors.Is(err, errdetail.ErrPermissionDenied):
		status = http.StatusForbidden
		title = errdetail.ErrPermissionDenied.Error()
		code = statusPermissionDenied

	case errors.Is(err, errdetail.ErrResourceExhausted):
		status = http.StatusTooManyRequests
		title = errdetail.ErrResourceExhausted.Error()
		code = statusResourceExhausted

	case errors.Is(err, errdetail.ErrDataCorrupted):
		status = http.StatusInternalServerError
		title = errdetail.ErrDataCorrupted.Error()
		code = statusDataCorrupted

	case errors.Is(err, errdetail.ErrInternal):
		status = http.StatusInternalServerError
		title = errdetail.ErrInternal.Error()
		code = statusInternal

	case errors.Is(err, errdetail.ErrNotImplemented):
		status = http.StatusNotImplemented
		title = errdetail.ErrNotImplemented.Error()
		code = statusNotImplemented

	case errors.Is(err, errdetail.ErrUnavailable):
		status = http.StatusServiceUnavailable
		title = errdetail.ErrUnavailable.Error()
		code = statusUnavailable

	case errors.Is(err, context.DeadlineExceeded), errors.Is(err, errdetail.ErrDeadlineExceeded):
		status = http.StatusGatewayTimeout
		title = errdetail.ErrDeadlineExceeded.Error()
		code = statusDeadlineExceeded

	case errors.Is(err, context.Canceled), errors.Is(err, errdetail.ErrCancelled):
		status = http.StatusGatewayTimeout
		title = errdetail.ErrCancelled.Error()
		code = statusCancelled

	default:
		status = http.StatusInternalServerError
		title = "unknown"
		code = statusUnknown
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
