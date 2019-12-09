package app

import (
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
)

const (
	minEmailLength = 3
	maxEmailLength = 320
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

func CheckRequiredMessageParams(messageParams bus.MessageParams, keys []string) error {
	for _, key := range keys {
		if messageParams[key] == nil {
			return Error(
				RequiredParamNotFound,
				"Message doesn't contain all required parameters",
			)
		}
	}

	return nil
}
