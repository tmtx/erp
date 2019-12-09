package validator

type Message string
type Messages map[string][]Message
type Group map[string][]Validator

type Validator interface {
	Validate() (bool, Message)
}

func ValidateGroup(g Group) (bool, Messages) {
	messages := map[string][]Message{}
	isValid := true
	for key, validators := range g {
		for _, validator := range validators {
			if valid, msg := validator.Validate(); !valid {
				messages[key] = append(messages[key], msg)
				isValid = false
			}
		}
	}

	return isValid, messages
}

func MergeMessages(groupedMessagesParams ...Messages) Messages {
	totalMessages := Messages{}
	for _, groupedMessages := range groupedMessagesParams {
		for key, messages := range groupedMessages {
			totalMessages[key] = append(totalMessages[key], messages...)
		}
	}

	return totalMessages
}
