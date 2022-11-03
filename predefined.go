// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package errdetail

type predefined string

// Error is the `error` interface implementation for the `predefined` type.
func (e predefined) Error() string {
	return string(e)
}

const (
	// ErrInvalidArgument - an invalid argument provided.
	// Error message and error details should provide more information.
	ErrInvalidArgument predefined = "invalid argument"
	// ErrFailedPrecondition - request can not be executed in the current system state,
	// such as deleting a non-empty directory.
	ErrFailedPrecondition predefined = "precondition failed"
	// ErrOutOfRange - an invalid range provided.
	ErrOutOfRange predefined = "out of range"
	// ErrUnauthenticated - not authenticated request due to missing, invalid or expired credentials.
	ErrUnauthenticated predefined = "unauthenticated"
	// ErrPermissionDenied - client does not have sufficient permission.
	ErrPermissionDenied predefined = "permission denied"
	// ErrNotFound - a specified resource is not found.
	ErrNotFound predefined = "not found"
	// ErrAborted - concurrency conflict, such as read-modify-write conflict.
	ErrAborted predefined = "aborted"
	// ErrAlreadyExists - the resource that a client tried to create already exists.
	ErrAlreadyExists predefined = "already exists"
	// ErrRemoved - a specified resource is no longer available at the origin server
	// and that this condition is likely to be permanent.
	ErrRemoved predefined = "removed"
	// ErrResourceExhausted - either out of resource quota or reaching rate limiting.
	// Error message and error details should provide more information.
	ErrResourceExhausted predefined = "resource exhausted"
	// ErrDataCorrupted - unrecoverable data loss or data corruption.
	ErrDataCorrupted predefined = "data corrupted"
	// ErrInternal - internal server error. Typically, a server bug.
	ErrInternal predefined = "internal"
	// ErrNotImplemented -  API method is not implemented by the server.
	ErrNotImplemented predefined = "not implemented"
	// ErrUnavailable - service unavailable. Typically, the server is down.
	ErrUnavailable predefined = "unavailable"
	// ErrDeadlineExceeded - request deadline exceeded (i.e. requested deadline is not enough
	// for the server to process the request).
	ErrDeadlineExceeded predefined = "deadline exceeded"
	// ErrCancelled - request cancelled by its creator.
	ErrCancelled predefined = "cancelled"
)

// NewInvalidArgument - sugar wrapper for ErrInvalidArgument.
func NewInvalidArgument(msg string, details ...Detail) error {
	return Wrap(ErrInvalidArgument, msg, details...)
}

// NewFailedPrecondition - sugar wrapper for ErrFailedPrecondition.
func NewFailedPrecondition(msg string, details ...Detail) error {
	return Wrap(ErrFailedPrecondition, msg, details...)
}

// NewOutOfRange - sugar wrapper for ErrOutOfRange.
func NewOutOfRange(msg string, details ...Detail) error {
	return Wrap(ErrOutOfRange, msg, details...)
}

// NewUnauthenticated - sugar wrapper for ErrUnauthenticated.
func NewUnauthenticated(msg string, details ...Detail) error {
	return Wrap(ErrUnauthenticated, msg, details...)
}

// NewPermissionDenied - sugar wrapper for ErrPermissionDenied.
func NewPermissionDenied(msg string, details ...Detail) error {
	return Wrap(ErrPermissionDenied, msg, details...)
}

// NewNotFound - sugar wrapper for ErrNotFound.
func NewNotFound(msg string, details ...Detail) error {
	return Wrap(ErrNotFound, msg, details...)
}

// NewAborted - sugar wrapper for ErrAborted.
func NewAborted(msg string, details ...Detail) error {
	return Wrap(ErrAborted, msg, details...)
}

// NewAlreadyExists - sugar wrapper for ErrAlreadyExists.
func NewAlreadyExists(msg string, details ...Detail) error {
	return Wrap(ErrAlreadyExists, msg, details...)
}

// NewRemoved - sugar wrapper for ErrRemoved.
func NewRemoved(msg string, details ...Detail) error {
	return Wrap(ErrRemoved, msg, details...)
}

// NewResourceExhausted - sugar wrapper for ErrResourceExhausted.
func NewResourceExhausted(msg string, details ...Detail) error {
	return Wrap(ErrResourceExhausted, msg, details...)
}

// NewDataCorrupted - sugar wrapper for ErrDataCorrupted.
func NewDataCorrupted(msg string, details ...Detail) error {
	return Wrap(ErrDataCorrupted, msg, details...)
}

// NewInternal - sugar wrapper for ErrInternal.
func NewInternal(msg string, details ...Detail) error {
	return Wrap(ErrInternal, msg, details...)
}

// NewNotImplemented - sugar wrapper for ErrNotImplemented.
func NewNotImplemented(msg string, details ...Detail) error {
	return Wrap(ErrNotImplemented, msg, details...)
}

// NewUnavailable - sugar wrapper for ErrUnavailable.
func NewUnavailable(msg string, details ...Detail) error {
	return Wrap(ErrUnavailable, msg, details...)
}

// NewDeadlineExceeded - sugar wrapper for ErrDeadlineExceeded.
func NewDeadlineExceeded(msg string, details ...Detail) error {
	return Wrap(ErrDeadlineExceeded, msg, details...)
}

// NewCancelled - sugar wrapper for ErrCancelled.
func NewCancelled(msg string, details ...Detail) error {
	return Wrap(ErrCancelled, msg, details...)
}
