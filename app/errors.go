// TODO: move to different package
package app

import (
	"fmt"
)

type ErrorType string

type ErrorHandler interface {
	Handle(err error)
}

type TypedError struct {
	Message string
	Type    ErrorType
}

const (
	UndefinedError           ErrorType = "undefined_error"
	RequiredParamNotFound    ErrorType = "required_param_not_found"
	ParamTypeError           ErrorType = "param_type_error"
	UnableToRestoreAggregate ErrorType = "aggregate_restore_error"
)

func (err TypedError) Error() string {
	if err.Message != "" {
		return fmt.Sprintf("%s: %s", err.Type, err.Message)
	}

	switch err.Type {
	case UndefinedError:
		err.Message = "Undefined error"
	case RequiredParamNotFound:
		err.Message = "Required parameter not found"
	case ParamTypeError:
		err.Message = "Parameter is of wrong type"
	case UnableToRestoreAggregate:
		err.Message = "Unable to restore aggregate"
	}

	return fmt.Sprintf("%s: %s", err.Type, err.Message)
}
