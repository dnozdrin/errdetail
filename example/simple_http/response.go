// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package simple_http

// ResponseCode represents codes for exposing publicly instead of error messages.
type ResponseCode string

const (
	statusInvalidArgument    ResponseCode = "INVALID_ARGUMENT"
	statusFailedPrecondition ResponseCode = "FAILED_PRECONDITION"
	statusOutOfRange         ResponseCode = "OUT_OF_RANGE"
	statusUnauthenticated    ResponseCode = "UNAUTHENTICATED"
	statusPermissionDenied   ResponseCode = "PERMISSION_DENIED"
	statusNotFound           ResponseCode = "NOT_FOUND"
	statusAborted            ResponseCode = "ABORTED"
	statusAlreadyExists      ResponseCode = "ALREADY_EXISTS"
	statusRemoved            ResponseCode = "REMOVED"
	statusResourceExhausted  ResponseCode = "RESOURCE_EXHAUSTED"
	statusDataCorrupted      ResponseCode = "DATA_CORRUPTED"
	statusInternal           ResponseCode = "INTERNAL"
	statusUnknown            ResponseCode = "UNKNOWN"
	statusNotImplemented     ResponseCode = "NOT_IMPLEMENTED"
	statusUnavailable        ResponseCode = "UNAVAILABLE"
	statusDeadlineExceeded   ResponseCode = "DEADLINE_EXCEEDED"
	statusCancelled          ResponseCode = "CANCELLED"
)
