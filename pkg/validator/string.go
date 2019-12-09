package validator

import (
	"fmt"
	"strings"
)

type StringNonEmpty struct{ Value string }
type StringLengthValid struct {
	Value     string
	MinLength int
	MaxLength int
}
type StringsEqual struct {
	Value1 string
	Value2 string
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

func (v StringsEqual) Validate() (bool, Message) {
	isValid := strings.TrimSpace(v.Value1) == strings.TrimSpace(v.Value2)

	message := Message("")
	if !isValid {
		message = Message("String value mismatch")
	}

	return isValid, message
}
