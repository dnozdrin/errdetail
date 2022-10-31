// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package http_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dnozdrin/errdetail"
	. "github.com/dnozdrin/errdetail/examples/http"
)

//go:embed testdata
var testdata embed.FS

func TestNewErrorResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		file string
	}{
		// todo: double wrapped error
		{
			name: "not_found/simple",
			err:  errdetail.ErrNotFound,
			file: "not_found_simple",
		},
		{
			name: "not_found/full",
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
		{
			name: "not_found/wrapped",
			err:  fmt.Errorf("test error: %w", errdetail.Wrap(errdetail.ErrNotFound, "not found")),
			file: "not_found_wrapped",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewErrorResponse(tt.err)

			actual, err := json.Marshal(got)
			require.NoError(t, err)

			expected, err := testdata.ReadFile("testdata/" + tt.file + ".json")
			require.NoError(t, err)

			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
