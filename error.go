// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package errdetail

import (
	"errors"
	"fmt"
)

// New allows to create unspecified errors with details provided.
func New(msg string, details ...Detail) error {
	return &wrapper{
		msg:        msg,
		underlying: nil,
		details:    filter(details),
	}
}

// Wrap allows to wrap errors and add details to them.
func Wrap(err error, msg string, details ...Detail) error {
	if err != nil {
		wrapped := &wrapper{underlying: err}

		if msg != "" {
			wrapped.msg = fmt.Sprintf("%s: %s", msg, err.Error())
		} else {
			wrapped.msg = err.Error()
		}

		wrapped.details = filter(ExtractDetails(err), details)

		return wrapped
	}

	return nil
}

func filter(src ...[]Detail) []Detail {
	var total int
	for i := range src {
		total += len(src[i])
	}

	result := make([]Detail, total)

	var n int
	for i := range src { //nolint:wsl // false positive
		for j := range src[i] {
			if src[i][j].filled {
				result[n] = src[i][j]
				n++
			}
		}
	}

	if n == 0 {
		return nil
	}

	return result[:n]
}

type wrapper struct {
	msg        string
	underlying error
	details    []Detail
}

// Error returns error message.
func (err *wrapper) Error() string {
	return err.msg
}

// Is reports whether any error in the wrapped error's chain matches target.
func (err *wrapper) Is(target error) bool {
	return errors.Is(err.underlying, target)
}

// Unwrap returns the underlying error that was wrapped.
func (err *wrapper) Unwrap() error {
	return err.underlying
}

// Details returns the wrapped error's details.
func (err *wrapper) Details() []Detail {
	return err.details
}
