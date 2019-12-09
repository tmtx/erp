package bus

import (
	"fmt"
	"testing"
	"time"

	"github.com/tmtx/erp/pkg/validator"
)

type testErrorHandler struct {
	err      error
	messages validator.Messages
}

func (eh *testErrorHandler) Handle(err error, messages validator.Messages) {
	eh.err = err
	eh.messages = messages
}

// TestErrorHandle tests if error is handled
func TestErrorHandle(t *testing.T) {
	var testErrorHandleMessageKey MessageKey = "test_error_handle_message_key"
	expectedErrorMessage := "1234"

	v := validator.StringNonEmpty{
		Value: "",
	}

	vg := validator.Group{
		"test": []validator.Validator{v},
	}
	eh := testErrorHandler{}
	messageBus := NewBasicMessageBus(&eh)
	messageBus.Subscribe(
		testErrorHandleMessageKey,
		func(p MessageParams) (validator.Messages, error) {
			isValid, messages := validator.ValidateGroup(vg)
			if isValid {
				t.Errorf("Expected validation to fail")
			}
			return messages, fmt.Errorf(expectedErrorMessage)
		},
	)
	messageBus.Listen()

	messageBus.Dispatch(Message{testErrorHandleMessageKey, MessageParams{}, CommandMessage})
	time.Sleep(time.Millisecond * 5)

	if eh.err.Error() != expectedErrorMessage {
		t.Errorf("Expected: '%s',  got: '%s'", expectedErrorMessage, eh.err.Error())
	}
}

// TestDispatch tests if multiple dispatch calls work properly
func TestDispatch(t *testing.T) {
	var testDispatchMessageKey MessageKey = "test_trigger_message_key"
	expectedResult := 5
	result := 0

	messageBus := NewBasicMessageBus(&testErrorHandler{})
	messageBus.Subscribe(
		testDispatchMessageKey,
		func(p MessageParams) (validator.Messages, error) {
			result++
			return nil, nil
		},
	)
	messageBus.Listen()

	m := Message{testDispatchMessageKey, MessageParams{}, CommandMessage}
	messageBus.Dispatch(m)
	messageBus.Dispatch(m)
	messageBus.Dispatch(m)
	messageBus.Dispatch(m)
	messageBus.Dispatch(m)
	time.Sleep(time.Millisecond * 5)

	if result != expectedResult {
		t.Errorf("Expected: '%d',  got: '%d'", expectedResult, result)
	}
}
