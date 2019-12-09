package guests

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/validator"
)

const (
	minNameLength = 3
	maxNameLength = 200
)

func ValidateCreateGuest(params app.CreateGuestParams) (bool, validator.Messages) {
	nameValidators := []validator.Validator{
		validator.StringLengthValid{
			Value:     params.Name,
			MinLength: minNameLength,
			MaxLength: maxNameLength,
		},
		validator.StringNonEmpty{
			Value: params.Name,
		},
	}
	emailValidators := app.EmailValidators(params.Email)

	return validator.ValidateGroup(validator.Group{
		"name":  nameValidators,
		"email": emailValidators,
	})
}
