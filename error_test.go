// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package errdetail_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/dnozdrin/errdetail"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		msg     string
		details []Detail
	}

	tests := map[string]struct {
		args        args
		wantDetails []Detail
	}{
		"full": {
			args: args{
				msg: "test message 1",
				details: []Detail{
					NewDetail(WithCode("dummy_code_1")),
					NewDetail(WithDescription("dummy description_1")),
				},
			},
			wantDetails: []Detail{
				NewDetail(WithCode("dummy_code_1")),
				NewDetail(WithDescription("dummy description_1")),
			},
		},
		"empty_message": {
			args: args{
				msg: "",
				details: []Detail{
					NewDetail(WithCode("dummy_code_2")),
					NewDetail(WithDescription("dummy description_2")),
				},
			},
			wantDetails: []Detail{
				NewDetail(WithCode("dummy_code_2")),
				NewDetail(WithDescription("dummy description_2")),
			},
		},
		"partially_filled_details": {
			args: args{
				msg: "test message 2",
				details: []Detail{
					NewDetail(),
					NewDetail(WithCode("")),
					NewDetail(WithCode("dummy_code_3")),
					NewDetail(
						WithCode(""),
						WithDescription(""),
					),
				},
			},
			wantDetails: []Detail{
				NewDetail(WithCode("dummy_code_3")),
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := New(tt.args.msg, tt.args.details...)
			require.Error(t, err)

			assert.Equal(t, tt.args.details, tt.args.details)
			assert.Equal(t, tt.wantDetails, ExtractDetails(err))
			assert.EqualError(t, err, tt.args.msg)
		})
	}
}

func TestWrapNil(t *testing.T) {
	t.Parallel()

	type args struct {
		err     error
		msg     string
		details []Detail
	}

	tests := map[string]struct {
		args args
	}{
		"all_empty": {
			args: args{
				err:     nil,
				msg:     "",
				details: nil,
			},
		},
		"with_message": {
			args: args{
				err:     nil,
				msg:     "dummy message",
				details: nil,
			},
		},
		"with_details": {
			args: args{
				err: nil,
				msg: "",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
			},
		},
		"with_empty_details": {
			args: args{
				err:     nil,
				msg:     "",
				details: []Detail{},
			},
		},
		"with_not_filled_details": {
			args: args{
				err: nil,
				msg: "",
				details: []Detail{
					NewDetail(),
					NewDetail(WithCode("")),
					NewDetail(
						WithCode(""),
						WithDescription(""),
					),
				},
			},
		},
		"with_message_and_details": {
			args: args{
				err: nil,
				msg: "dummy message",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := Wrap(tt.args.err, tt.args.msg, tt.args.details...)
			assert.NoError(t, err)
		})
	}
}

func TestWrap(t *testing.T) {
	t.Parallel()

	type args struct {
		err     error
		msg     string
		details []Detail
	}

	type want struct {
		msg     string
		details []Detail
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"no_message/nil_details": {
			args: args{
				err:     assert.AnError,
				msg:     "",
				details: nil,
			},
			want: want{
				details: nil,
				msg:     assert.AnError.Error(),
			},
		},
		"no_message/with_details": {
			args: args{
				err: assert.AnError,
				msg: "",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
			},
			want: want{
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
				msg: assert.AnError.Error(),
			},
		},
		"nil_details": {
			args: args{
				err:     assert.AnError,
				msg:     "dummy message",
				details: nil,
			},
			want: want{
				details: nil,
				msg:     "dummy message: " + assert.AnError.Error(),
			},
		},
		"empty_details": {
			args: args{
				err:     assert.AnError,
				msg:     "dummy message",
				details: []Detail{},
			},
			want: want{
				details: nil,
				msg:     "dummy message: " + assert.AnError.Error(),
			},
		},
		"with_not_filled_details": {
			args: args{
				err: assert.AnError,
				msg: "dummy message",
				details: []Detail{
					NewDetail(),
					NewDetail(WithCode("")),
					NewDetail(
						WithCode(""),
						WithDescription(""),
					),
				},
			},
			want: want{
				details: nil,
				msg:     "dummy message: " + assert.AnError.Error(),
			},
		},
		"with_partially_filled_details": {
			args: args{
				err: assert.AnError,
				msg: "dummy message",
				details: []Detail{
					NewDetail(),
					NewDetail(WithCode("")),
					NewDetail(WithCode("dummy_code")),
					NewDetail(
						WithCode(""),
						WithDescription(""),
					),
				},
			},
			want: want{
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
				},
				msg: "dummy message: " + assert.AnError.Error(),
			},
		},
		"with_message_and_details": {
			args: args{
				err: assert.AnError,
				msg: "dummy message",
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
			},
			want: want{
				details: []Detail{
					NewDetail(WithCode("dummy_code")),
					NewDetail(WithDescription("dummy description")),
				},
				msg: "dummy message: " + assert.AnError.Error(),
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := Wrap(tt.args.err, tt.args.msg, tt.args.details...)
			require.Error(t, err)

			assert.Equal(t, tt.want.msg, err.Error())
			assert.Equal(t, tt.want.details, ExtractDetails(err))
			assert.Equal(t, tt.args.err, errors.Unwrap(err))
			assert.ErrorIs(t, err, tt.args.err)

			var dErr detailedError
			if assert.Truef(t, errors.As(err, &dErr), "invalid type") {
				assert.Equal(t, tt.want.details, dErr.Details())
				assert.Equal(t, tt.args.err, dErr.Unwrap())
				assert.True(t, dErr.Is(tt.args.err))
			}
		})
	}
}

type detailedError interface {
	Is(error) bool
	Unwrap() error
	Details() []Detail
}

func TestChainedWrap(t *testing.T) {
	t.Parallel()

	t.Run("wrap_error_previously_wrapped_by_std_lib", func(t *testing.T) {
		t.Parallel()

		err1 := assert.AnError
		err2 := fmt.Errorf("second error: %w", err1)
		err3 := fmt.Errorf("third error: %w", err2)

		details := []Detail{
			NewDetail(WithField("dummy_field")),
			NewDetail(WithCode("dummy_code")),
			NewDetail(WithDescription("dummy description")),
		}

		err4 := Wrap(err3, "fourth", details...)
		require.Error(t, err4)

		assert.ErrorIs(t, err4, err1)

		var dErr detailedError
		if assert.Truef(t, errors.As(err4, &dErr), "invalid type") {
			assert.Equal(t, details, dErr.Details())
			assert.Equal(t, err3, dErr.Unwrap())
			assert.True(t, dErr.Is(err1))
			assert.True(t, dErr.Is(err2))
			assert.True(t, dErr.Is(err3))
		}
	})

	t.Run("wrap_error_by_std_lib_previously_wrapped_by_this_lib", func(t *testing.T) {
		t.Parallel()

		err1 := assert.AnError
		details := []Detail{
			NewDetail(WithField("dummy_field")),
			NewDetail(WithCode("dummy_code")),
			NewDetail(WithDescription("dummy description")),
		}

		err2 := Wrap(err1, "second", details...)
		require.Error(t, err2)

		err3 := fmt.Errorf("third error: %w", err2)

		assert.ErrorIs(t, err3, err2)
		assert.ErrorIs(t, err3, err1)
		assert.ErrorAs(t, err2, new(detailedError))
		assert.ErrorAs(t, err3, new(detailedError))
	})

	t.Run("wrap_multiple_times_by_this_lib", func(t *testing.T) {
		t.Parallel()

		err1 := assert.AnError
		details := []Detail{
			NewDetail(WithField("dummy_field")),
			NewDetail(WithCode("dummy_code")),
			NewDetail(WithDescription("dummy description")),
		}

		err2 := Wrap(err1, "second", details...)
		require.Error(t, err2)

		err3 := Wrap(err2, "third", details...)
		require.Error(t, err3)

		assert.ErrorIs(t, err3, err2)
		assert.ErrorIs(t, err3, err1)
		assert.ErrorIs(t, err2, err1)
	})
}
