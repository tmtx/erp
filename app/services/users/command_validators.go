package users

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type passwordValidator struct {
	Password       []byte
	HashedPassword []byte
}

func ValidateLoginUser(u *aggregates.User, p app.LoginUserParams) (bool, validator.Messages) {
	if u == nil {
		return false, validator.Messages{
			"email": []validator.Message{
				validator.Message("User not found"),
			},
		}
	}

	emailValidators := app.EmailValidators(p.Email)
	emailValidators = append(
		emailValidators,
		validator.StringsEqual{
			Value1: u.Email,
			Value2: p.Email,
		},
	)

	g := validator.Group{
		"email": emailValidators,
		"password": []validator.Validator{
			passwordValidator{
				HashedPassword: []byte(u.HashedPassword),
				Password:       []byte(p.Password),
			},
		},
	}

	return validator.ValidateGroup(g)
}

func ValidateUserInfo(p app.UpdateUserInfoParams) (bool, validator.Messages) {
	emailValidators := app.EmailValidators(p.Email)

	g := validator.Group{
		"email":  emailValidators,
		"userId": []validator.Validator{validator.NonNilValidator{p.UserId}},
	}

	return validator.ValidateGroup(g)
}

func (v passwordValidator) Validate() (bool, validator.Message) {
	err := bcrypt.CompareHashAndPassword(v.HashedPassword, v.Password)
	if err != nil {
		return false, validator.Message("Password invalid")
	}

	return true, validator.Message("")
}
