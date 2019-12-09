package validator

type NonNilValidator struct {
	Value interface{}
}

func (v NonNilValidator) Validate() (bool, Message) {
	if v.Value == nil {
		return false, Message("String is empty")
	}

	return true, ""
}
