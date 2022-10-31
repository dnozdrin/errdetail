// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package http

// ResponseStatus represents codes for exposing publicly instead of error messages.
type ResponseStatus string

const (
	statusInvalidArgument    ResponseStatus = "INVALID_ARGUMENT"
	statusFailedPrecondition ResponseStatus = "FAILED_PRECONDITION"
	statusOutOfRange         ResponseStatus = "OUT_OF_RANGE"
	statusUnauthenticated    ResponseStatus = "UNAUTHENTICATED"
	statusNotFound           ResponseStatus = "NOT_FOUND"
	statusAborted            ResponseStatus = "ABORTED"
	statusAlreadyExists      ResponseStatus = "ALREADY_EXISTS"
	statusRemoved            ResponseStatus = "REMOVED"
	statusPermissionDenied   ResponseStatus = "PERMISSION_DENIED"
	statusResourceExhausted  ResponseStatus = "RESOURCE_EXHAUSTED"
	statusDataCorrupted      ResponseStatus = "DATA_CORRUPTED"
	statusInternal           ResponseStatus = "INTERNAL"
	statusUnknown            ResponseStatus = "UNKNOWN"
	statusNotImplemented     ResponseStatus = "NOT_IMPLEMENTED"
	statusUnavailable        ResponseStatus = "UNAVAILABLE"
	statusDeadlineExceeded   ResponseStatus = "DEADLINE_EXCEEDED"
	statusCancelled          ResponseStatus = "CANCELLED"
)
