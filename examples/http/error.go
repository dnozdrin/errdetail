// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/dnozdrin/errdetail"
)

// ErrorResponse represents an HTTP server response in case of an error.
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error is expected to be filled only via NewErrorResponse.
// Intentionally omits error's message presentation, instead ResponseStatus should be
// exposed publicly.
type Error struct {
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
	Details []ErrorDetail  `json:"details,omitempty"`
}

// ErrorDetail is expected to be filled only via NewErrorResponse.
type ErrorDetail struct {
	Domain      string `json:"domain,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
	Code        string `json:"code,omitempty"`
}

// NewErrorResponse creates a ErrorResponse with properly filled fields.
func NewErrorResponse(err error) ErrorResponse {
	code, status := resolveCodeAndStatus(err)

	return ErrorResponse{
		Error: Error{
			Code:    code,
			Status:  status,
			Details: toErrorDetails(errdetail.ExtractDetails(err)),
		},
	}
}

func resolveCodeAndStatus(err error) (int, ResponseStatus) {
	switch {
	case errors.Is(err, errdetail.ErrInvalidArgument):
		return http.StatusBadRequest, statusInvalidArgument

	case errors.Is(err, errdetail.ErrFailedPrecondition):
		return http.StatusBadRequest, statusFailedPrecondition

	case errors.Is(err, errdetail.ErrOutOfRange):
		return http.StatusBadRequest, statusOutOfRange

	case errors.Is(err, errdetail.ErrUnauthenticated):
		return http.StatusUnauthorized, statusUnauthenticated

	case errors.Is(err, errdetail.ErrNotFound):
		return http.StatusNotFound, statusNotFound

	case errors.Is(err, errdetail.ErrAborted):
		return http.StatusConflict, statusAborted

	case errors.Is(err, errdetail.ErrAlreadyExists):
		return http.StatusConflict, statusAlreadyExists

	case errors.Is(err, errdetail.ErrRemoved):
		return http.StatusGone, statusRemoved

	case errors.Is(err, errdetail.ErrPermissionDenied):
		return http.StatusForbidden, statusPermissionDenied

	case errors.Is(err, errdetail.ErrResourceExhausted):
		return http.StatusTooManyRequests, statusResourceExhausted

	case errors.Is(err, errdetail.ErrDataCorrupted):
		return http.StatusInternalServerError, statusDataCorrupted

	case errors.Is(err, errdetail.ErrInternal):
		return http.StatusInternalServerError, statusInternal

	case errors.Is(err, errdetail.ErrNotImplemented):
		return http.StatusNotImplemented, statusNotImplemented

	case errors.Is(err, errdetail.ErrUnavailable):
		return http.StatusServiceUnavailable, statusUnavailable

	case errors.Is(err, context.DeadlineExceeded), errors.Is(err, errdetail.ErrDeadlineExceeded):
		return http.StatusGatewayTimeout, statusDeadlineExceeded

	case errors.Is(err, context.Canceled), errors.Is(err, errdetail.ErrCancelled):
		return http.StatusGatewayTimeout, statusCancelled

	default:
		return http.StatusInternalServerError, statusUnknown
	}
}

func toErrorDetails(input []errdetail.Detail) []ErrorDetail {
	if len(input) == 0 {
		return nil
	}

	details := make([]ErrorDetail, len(input))
	for i := range input {
		details[i] = ErrorDetail{
			Domain:      input[i].Domain(),
			Reason:      input[i].Reason(),
			Field:       input[i].Field(),
			Description: input[i].Description(),
			Code:        input[i].Code(),
		}
	}

	return details
}
