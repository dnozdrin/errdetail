// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package errdetail_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/dnozdrin/errdetail"
)

func TestDetail(t *testing.T) {
	t.Parallel()

	type args struct {
		opts []Option
	}

	type values struct {
		field       string
		description string
		code        string
		domain      string
		reason      string
	}

	tests := []struct {
		name   string
		args   args
		values values
	}{
		{
			name: "with_all_options",
			args: args{
				opts: []Option{
					WithField("field_1"),
					WithDescription("description_1"),
					WithCode("code_1"),
					WithDomain("domain_1"),
					WithReason("reason_1"),
				},
			},
			values: values{
				field:       "field_1",
				description: "description_1",
				code:        "code_1",
				domain:      "domain_1",
				reason:      "reason_1",
			},
		},
		{
			name: "empty_options",
			args: args{
				opts: nil,
			},
			values: values{},
		},
		{
			name: "nil_options",
			args: args{
				opts: nil,
			},
			values: values{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			detail := NewDetail(tt.args.opts...)

			assert.Equal(t, tt.values.field, detail.Field())
			assert.Equal(t, tt.values.description, detail.Description())
			assert.Equal(t, tt.values.code, detail.Code())
			assert.Equal(t, tt.values.domain, detail.Domain())
			assert.Equal(t, tt.values.reason, detail.Reason())
		})
	}
}

func TestExtractDetails(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want []Detail
	}{
		{
			name: "no_details_in_error",
			args: args{
				err: assert.AnError,
			},
			want: nil,
		},
		{
			name: "no_error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "with_details",
			args: args{
				err: New("test_error",
					NewDetail(
						WithDomain("dummy_domain_1"),
						WithCode("dummy_code_1"),
					),
					NewDetail(
						WithDescription("dummy_description_1"),
						WithField("dummy_field_1"),
						WithReason("dummy_reason_1"),
					),
				),
			},
			want: []Detail{
				NewDetail(
					WithDomain("dummy_domain_1"),
					WithCode("dummy_code_1"),
				),
				NewDetail(
					WithDescription("dummy_description_1"),
					WithField("dummy_field_1"),
					WithReason("dummy_reason_1"),
				),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.NotPanics(t, func() {
				got := ExtractDetails(tt.args.err)
				assert.Equalf(t, tt.want, got, "ExtractDetails(%v)", tt.args.err)
			})
		})
	}
}
