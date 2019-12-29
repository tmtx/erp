package app

import (
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
)

const (
	minEmailLength = 3
	maxEmailLength = 320
	minNameLength  = 3
	maxNameLength  = 100
)

var allowedEmailDomains = [...]string{
	"gmail.com",
	"karlis.dev",
}

func EmailValidators(val string) []validator.Validator {
	return []validator.Validator{
		validator.StringLengthValid{
			Value:     val,
			MinLength: minEmailLength,
			MaxLength: maxEmailLength,
		},
		validator.StringNonEmpty{
			Value: val,
		},
		validator.EmailDomainAllowed{
			Value:          val,
			AllowedDomains: allowedEmailDomains[:],
		},
	}
}

func NameValidators(val string) []validator.Validator {
	return []validator.Validator{
		validator.StringLengthValid{
			Value:     val,
			MinLength: minNameLength,
			MaxLength: maxNameLength,
		},
		validator.StringNonEmpty{
			Value: val,
		},
	}
}

func CheckRequiredMessageParams(messageParams bus.MessageParams, keys []string) error {
	for _, key := range keys {
		if messageParams[key] == nil {
			return TypedError{Type: RequiredParamNotFound}
		}
	}

	return nil
}
