package validator

import (
	"fmt"
)

type StringNonEmpty struct{ Value string }
type StringLengthValid struct {
	Value     string
	MinLength int
	MaxLength int
}

func (v StringNonEmpty) Validate() (bool, Message) {
	if v.Value == "" {
		return false, Message("String is empty")
	}

	return true, ""
}

func (v StringLengthValid) Validate() (bool, Message) {
	if len(v.Value) < v.MinLength {
		return false, Message(fmt.Sprintf("Min string length is %d", v.MinLength))
	}
	if len(v.Value) > v.MaxLength {
		return false, Message(fmt.Sprintf("Max string length is %d", v.MaxLength))
	}

	return true, ""
}
