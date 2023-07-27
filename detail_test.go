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
		meta        Meta
	}

	tests := map[string]struct {
		args   args
		values values
	}{
		"with_all_options": {
			args: args{
				opts: []Option{
					WithField("field_1"),
					WithDescription("description_1"),
					WithCode("code_1"),
					WithDomain("domain_1"),
					WithReason("reason_1"),
					WithMeta(Meta{
						"dummyField1": 1,
						"dummyField2": "dummyValue",
						"dummyField3": map[string]int{
							"1": 1,
							"2": 2,
						},
					}),
				},
			},
			values: values{
				field:       "field_1",
				description: "description_1",
				code:        "code_1",
				domain:      "domain_1",
				reason:      "reason_1",
				meta: Meta{
					"dummyField1": 1,
					"dummyField2": "dummyValue",
					"dummyField3": map[string]int{
						"1": 1,
						"2": 2,
					},
				},
			},
		},
		"empty_options": {
			args: args{
				opts: nil,
			},
			values: values{},
		},
		"nil_options": {
			args: args{
				opts: nil,
			},
			values: values{},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			detail := NewDetail(tt.args.opts...)

			assert.Equal(t, tt.values.field, detail.Field())
			assert.Equal(t, tt.values.description, detail.Description())
			assert.Equal(t, tt.values.code, detail.Code())
			assert.Equal(t, tt.values.domain, detail.Domain())
			assert.Equal(t, tt.values.reason, detail.Reason())
			assert.Equal(t, tt.values.meta, detail.Meta())
		})
	}
}

func TestExtractDetails(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
	}

	tests := map[string]struct {
		args args
		want []Detail
	}{
		"no_details_in_error": {
			args: args{
				err: assert.AnError,
			},
			want: nil,
		},
		"no_error": {
			args: args{
				err: nil,
			},
			want: nil,
		},
		"with_details": {
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
						WithMeta(Meta{"key1": "value1"}),
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
					WithMeta(Meta{"key1": "value1"}),
				),
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.NotPanics(t, func() {
				got := ExtractDetails(tt.args.err)
				assert.Equalf(t, tt.want, got, "ExtractDetails(%v)", tt.args.err)
			})
		})
	}
}
