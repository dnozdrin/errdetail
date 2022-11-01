// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package jsonapi_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/dnozdrin/errdetail/examples/json_api"

	"github.com/dnozdrin/errdetail"
)

func TestNewErrorResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		file string
	}{
		{
			name: "no_error",
			err:  nil,
			file: "no_error",
		},
		{
			name: "no_details",
			err:  assert.AnError,
			file: "no_error",
		},
		{
			name: "multiple_errors",
			err: errdetail.Wrap(
				assert.AnError,
				"let's do it again, %username%, everything is garbage",
				errdetail.NewDetail(
					errdetail.WithDomain("user.management"),
					errdetail.WithCode("invalid_email"),
					errdetail.WithDescription("Invalid email"),
					errdetail.WithField("user.email"),
					errdetail.WithReason("Email address \"lol@kek@cheburek\" is not valid."),
				),
				errdetail.NewDetail(
					errdetail.WithCode("admin_required"),
					errdetail.WithDescription("Permission denied"),
					errdetail.WithReason("Editing secret powers is not authorized on Sundays."),
				),
				errdetail.NewDetail(
					errdetail.WithCode("internal"),
					errdetail.WithDescription("The backend responded with an error"),
					errdetail.WithReason("Reputation service not responding after three requests."),
				),
			),
			file: "multiple_errors",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
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
