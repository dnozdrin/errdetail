// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package errdetail_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/dnozdrin/errdetail"
)

func TestErrorConstructors(t *testing.T) {
	t.Parallel()

	type args struct {
		msg     string
		details []Detail
	}
	tests := []struct {
		name        string
		constructor func(string, ...Detail) error
		err         error
	}{
		{
			name:        "invalid_argument",
			constructor: NewInvalidArgument,
			err:         ErrInvalidArgument,
		},
		{
			name:        "precondition_failed",
			constructor: NewFailedPrecondition,
			err:         ErrFailedPrecondition,
		},
		{
			name:        "out_of_range",
			constructor: NewOutOfRange,
			err:         ErrOutOfRange,
		},
		{
			name:        "unauthenticated",
			constructor: NewUnauthenticated,
			err:         ErrUnauthenticated,
		},
		{
			name:        "permission_denied",
			constructor: NewPermissionDenied,
			err:         ErrPermissionDenied,
		},
		{
			name:        "not_found",
			constructor: NewNotFound,
			err:         ErrNotFound,
		},
		{
			name:        "aborted",
			constructor: NewAborted,
			err:         ErrAborted,
		},
		{
			name:        "already_exists",
			constructor: NewAlreadyExists,
			err:         ErrAlreadyExists,
		},
		{
			name:        "removed",
			constructor: NewRemoved,
			err:         ErrRemoved,
		},
		{
			name:        "resource_exhausted",
			constructor: NewResourceExhausted,
			err:         ErrResourceExhausted,
		},
		{
			name:        "data_corrupted",
			constructor: NewDataCorrupted,
			err:         ErrDataCorrupted,
		},
		{
			name:        "internal",
			constructor: NewInternal,
			err:         ErrInternal,
		},
		{
			name:        "not_implemented",
			constructor: NewNotImplemented,
			err:         ErrNotImplemented,
		},
		{
			name:        "unavailable",
			constructor: NewUnavailable,
			err:         ErrUnavailable,
		},
		{
			name:        "deadline_exceeded",
			constructor: NewDeadlineExceeded,
			err:         ErrDeadlineExceeded,
		},
		{
			name:        "cancelled",
			constructor: NewCancelled,
			err:         ErrCancelled,
		},
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "full",
			args: args{
				msg: "dummy message",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
					NewDetail(WithField("dummy field")),
					NewDetail(WithDomain("dummy field")),
					NewDetail(WithReason("dummy field")),
				},
			},
		},
		{
			name: "message_only",
			args: args{
				msg:     "dummy message",
				details: nil,
			},
		},
		{
			name: "details_only",
			args: args{
				msg: "dummy message",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
					NewDetail(WithField("dummy field")),
					NewDetail(WithDomain("dummy field")),
					NewDetail(WithReason("dummy field")),
				},
			},
		},
		{
			name: "empty",
			args: args{
				msg:     "",
				details: nil,
			},
		},
	}
	for i := range tests {
		tt := tests[i]

		for j := range cases {
			tc := cases[j]

			t.Run(tt.name+"/"+tc.name, func(t *testing.T) {
				t.Parallel()

				details := make([]Detail, len(tc.args.details))
				copy(details, tc.args.details) // todo: avoid race conditions in code, not in tests!

				err := tt.constructor(tc.args.msg, details...)
				assert.Error(t, err)
				assert.Equal(t, tc.args.msg, err.Error())
				assert.Equal(t, tc.args.details, ExtractDetails(err))
				assert.Equal(t, tt.err, errors.Unwrap(err))
			})
		}
	}
}
