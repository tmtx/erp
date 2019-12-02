package guests

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/validator"
)

const (
	minNameLength  = 3
	maxNameLength  = 200
	minEmailLength = 3
	maxEmailLength = 320
)

var allowedEmailDomains = [...]string{
	"gmail.com",
	"karlis.dev",
}

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
	emailValidators := []validator.Validator{
		validator.StringLengthValid{
			Value:     params.Email,
			MinLength: minEmailLength,
			MaxLength: maxEmailLength,
		},
		validator.StringNonEmpty{
			Value: params.Email,
		},
		validator.EmailDomainAllowed{
			Value:          params.Email,
			AllowedDomains: allowedEmailDomains[:],
		},
	}

	return validator.ValidateGroup(validator.Group{
		"name":  nameValidators,
		"email": emailValidators,
	})
}
