package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dnozdrin/errdetail"
	. "github.com/dnozdrin/errdetail/examples/validation"
)

func TestValidateUser(t *testing.T) {
	t.Parallel()

	type args struct {
		user User
	}
	tests := []struct {
		name        string
		args        args
		assertErr   assert.ErrorAssertionFunc
		wantDetails []errdetail.Detail
	}{
		{
			name: "error/empty_name",
			args: args{
				user: User{
					Name:  "",
					Email: "dummy@test.com",
				},
			},
			wantDetails: []errdetail.Detail{
				errdetail.NewDetail(
					errdetail.WithField("name"),
					errdetail.WithCode("validation_required"),
					errdetail.WithDescription("cannot be blank"),
				),
			},
			assertErr: assert.Error,
		},
		{
			name: "error/too_long_name",
			args: args{
				user: User{
					Name:  "123456789012345678901234567890123456789012345678901",
					Email: "dummy@test.com",
				},
			},
			wantDetails: []errdetail.Detail{
				errdetail.NewDetail(
					errdetail.WithField("name"),
					errdetail.WithCode("validation_length_out_of_range"),
					errdetail.WithDescription("the length must be between 1 and 50"),
				),
			},
			assertErr: assert.Error,
		},
		{
			name: "error/empty_email",
			args: args{
				user: User{
					Name:  "test user",
					Email: "",
				},
			},
			wantDetails: []errdetail.Detail{
				errdetail.NewDetail(
					errdetail.WithField("email"),
					errdetail.WithCode("validation_required"),
					errdetail.WithDescription("cannot be blank"),
				),
			},
			assertErr: assert.Error,
		},
		{
			name: "error/invalid_email_format",
			args: args{
				user: User{
					Name:  "test user",
					Email: "@@@@invalid@@@@",
				},
			},
			wantDetails: []errdetail.Detail{
				errdetail.NewDetail(
					errdetail.WithField("email"),
					errdetail.WithCode("validation_is_email"),
					errdetail.WithDescription("must be a valid email address"),
				),
			},
			assertErr: assert.Error,
		},
		{
			name: "success",
			args: args{
				user: User{
					Name:  "test user",
					Email: "dummy@test.com",
				},
			},
			wantDetails: nil,
			assertErr:   assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateUser(tt.args.user)
			tt.assertErr(t, err)

			assert.Equal(t, tt.wantDetails, errdetail.ExtractDetails(err))
		})
	}
}
