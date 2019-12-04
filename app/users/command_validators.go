package users

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/validator"
)

// TODO: implement
func ValidateLoginUser(u *User, p app.LoginUserParams) (bool, *validator.Messages) {
	return true, nil
}
