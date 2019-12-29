package guests

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/pkg/validator"
)

func ValidateCreateGuest(params app.CreateGuestParams) (bool, validator.Messages) {
	nameValidators := app.NameValidators(params.Name)
	emailValidators := app.EmailValidators(params.Email)

	return validator.ValidateGroup(validator.Group{
		"name":  nameValidators,
		"email": emailValidators,
	})
}
