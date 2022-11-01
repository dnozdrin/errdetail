package validation

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/dnozdrin/errdetail"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ValidateUser(user User) error {
	err := ozzo.ValidateStruct(&user,
		ozzo.Field(&user.Name, ozzo.Required, ozzo.RuneLength(1, 50)),
		ozzo.Field(&user.Email, ozzo.Required, is.EmailFormat),
	)

	return toDetailedError(err, "user validation")
}

func toDetailedError(err error, msg string) error {
	if err == nil {
		return nil
	}

	if errs, ok := err.(ozzo.Errors); ok {
		details := make([]errdetail.Detail, 0, len(errs))
		for field, err := range errs {
			switch e := err.(type) {
			case ozzo.Errors:
				for subfield, subErr := range e {
					if ee, ok := subErr.(ozzo.Error); ok {
						details = append(details, toErrorDetail(field+"."+subfield, ee))
					}
				}
			case ozzo.Error:
				details = append(details, toErrorDetail(field, e))
			}
		}

		return errdetail.NewInvalidArgument(msg, details...)
	}

	return errdetail.NewInternal("validation internal failure", errdetail.NewDetail(
		errdetail.WithDescription(err.Error()),
		errdetail.WithReason("unexpected error"),
	))
}

func toErrorDetail(field string, err ozzo.Error) errdetail.Detail {
	return errdetail.NewDetail(
		errdetail.WithCode(err.Code()),
		errdetail.WithField(field),
		errdetail.WithDescription(err.Error()),
	)
}
