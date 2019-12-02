package app

import (
	"fmt"
)

type ErrorType string

type ErrorHandler interface {
	Handle(err error)
}

const (
	UndefinedError                  = "undefined_error"
	RequiredParamNotFound ErrorType = "required_param_not_found"
)

func Error(e ErrorType, msg string) error {
	return fmt.Errorf("%s: %s", e, msg)
}
