// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package simple_http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dnozdrin/errdetail"

	. "github.com/dnozdrin/errdetail/examples/simple_http"
)

func TestNewErrorResponse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err  error
		file string
	}{
		"no_error": {
			err:  nil,
			file: "no_error",
		},
		"bad_request": {
			err: errdetail.Wrap(
				errdetail.ErrInvalidArgument,
				"bad request",
				errdetail.NewDetail(
					errdetail.WithDomain("user.auth"),
					errdetail.WithCode("invalid_email"),
					errdetail.WithDescription("email validation failed"),
					errdetail.WithField("user.email"),
					errdetail.WithReason("an invalid character has been detected in the provided sequence"),
					errdetail.WithMeta(errdetail.Meta{
						"link": "https://example.com",
						"translations": map[string]string{
							"en": "Hello world!",
							"ua": "Привіт, світе!",
						},
					}),
				),
				errdetail.NewDetail(
					errdetail.WithDomain("user.auth"),
					errdetail.WithCode("invalid_password"),
					errdetail.WithDescription("password validation failed"),
					errdetail.WithField("user.password"),
					errdetail.WithReason("password is empty"),
				),
			),
			file: "bad_request",
		},
		"failed_precondition": {
			err: errdetail.Wrap(
				errdetail.ErrFailedPrecondition,
				"precondition failed",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "failed_precondition",
		},
		"out_of_range": {
			err: errdetail.Wrap(
				errdetail.ErrOutOfRange,
				"out of range",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "out_of_range",
		},
		"unauthenticated": {
			err: errdetail.Wrap(
				errdetail.ErrUnauthenticated,
				"unauthenticated",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "unauthenticated",
		},
		"permission_denied": {
			err: errdetail.Wrap(
				errdetail.ErrPermissionDenied,
				"permission denied",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "permission_denied",
		},
		"not_found/simple": {
			err:  errdetail.ErrNotFound,
			file: "not_found_simple",
		},
		"not_found/full": {
			err: errdetail.Wrap(
				errdetail.ErrNotFound,
				"not found full",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
					errdetail.WithDescription("dummy_description_1"),
					errdetail.WithField("dummy_field_1"),
					errdetail.WithReason("dummy_reason_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_2"),
					errdetail.WithCode("dummy_code_2"),
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_3"),
					errdetail.WithCode("dummy_code_3"),
					errdetail.WithDescription("dummy_description_3"),
					errdetail.WithField("dummy_field_3"),
					errdetail.WithReason("dummy_reason_3"),
				),
			),
			file: "not_found_full",
		},
		"not_found/wrapped": {
			err:  fmt.Errorf("test error: %w", errdetail.Wrap(errdetail.ErrNotFound, "not found")),
			file: "not_found_wrapped",
		},
		"not_found/double_wrapped": {
			err:  fmt.Errorf("test error: %w", errdetail.Wrap(fmt.Errorf("%w", errdetail.ErrNotFound), "not found")),
			file: "not_found_wrapped",
		},
		"aborted": {
			err: errdetail.Wrap(
				errdetail.ErrAborted,
				"aborted",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "aborted",
		},
		"already_exists": {
			err: errdetail.Wrap(
				errdetail.ErrAlreadyExists,
				"already exists",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "already_exists",
		},
		"removed": {
			err: errdetail.Wrap(
				errdetail.ErrRemoved,
				"removed",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "removed",
		},
		"resource_exhausted": {
			err: errdetail.Wrap(
				errdetail.ErrResourceExhausted,
				"resource exhausted",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "resource_exhausted",
		},
		"data_corrupted": {
			err: errdetail.Wrap(
				errdetail.ErrDataCorrupted,
				"data corrupted",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "data_corrupted",
		},
		"internal": {
			err: errdetail.Wrap(
				errdetail.ErrInternal,
				"internal",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "internal",
		},
		"unknown/wrapped": {
			err: errdetail.Wrap(
				assert.AnError,
				"dummy",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "unknown",
		},
		"unknown/new": {
			err: errdetail.New(
				"dummy",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "unknown",
		},
		"not_implemented": {
			err: errdetail.Wrap(
				errdetail.ErrNotImplemented,
				"not implemented",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "not_implemented",
		},
		"unavailable": {
			err: errdetail.Wrap(
				errdetail.ErrUnavailable,
				"unavailable",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "unavailable",
		},
		"deadline_exceeded/context": {
			err: errdetail.Wrap(
				context.DeadlineExceeded,
				"deadline exceeded",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "deadline_exceeded",
		},
		"deadline_exceeded/predefined": {
			err: errdetail.Wrap(
				errdetail.ErrDeadlineExceeded,
				"deadline exceeded",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "deadline_exceeded",
		},
		"cancelled/context": {
			err: errdetail.Wrap(
				context.Canceled,
				"cancelled",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "cancelled",
		},
		"cancelled/predefined": {
			err: errdetail.Wrap(
				errdetail.ErrCancelled,
				"cancelled",
				errdetail.NewDetail(
					errdetail.WithDomain("dummy_domain_1"),
					errdetail.WithCode("dummy_code_1"),
				),
				errdetail.NewDetail(
					errdetail.WithDescription("dummy_description_2"),
					errdetail.WithField("dummy_field_2"),
					errdetail.WithReason("dummy_reason_2"),
				),
			),
			file: "cancelled",
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := NewErrorResponse(tt.err)

			actual, err := json.Marshal(got)
			require.NoError(t, err)

			expected, err := os.ReadFile("testdata/" + tt.file + ".json")
			require.NoError(t, err)

			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
