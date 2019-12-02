package bus

import (
	"fmt"
	"testing"
	"time"
)

type testErrorHandler struct {
	handle func(err error)
}

func (eh *testErrorHandler) Handle(err error) {
	if eh.handle != nil {
		eh.handle(err)
	}
}

// TestErrorHandle tests if error is handled
func TestErrorHandle(t *testing.T) {
	var testErrorHandleMessageKey MessageKey = "test_error_handle_message_key"
	expectedResult := "1234"
	result := ""

	errorHandler := struct {
		testErrorHandler
	}{}
	errorHandler.handle = func(err error) {
		result = err.Error()
	}

	messageBus := NewBasicBus(&errorHandler)
	messageBus.Subscribe(testErrorHandleMessageKey, func(c Message) error {
		return fmt.Errorf(expectedResult)
	})
	messageBus.Listen()

	messageBus.Dispatch(Message{testErrorHandleMessageKey, MessageParams{}})
	time.Sleep(time.Millisecond * 5)

	if result != expectedResult {
		t.Errorf("Expected: '%s',  got: '%s'", expectedResult, result)
	}
}

// TestDispatch tests if multiple dispatch calls work properly
func TestDispatch(t *testing.T) {
	var testDispatchMessageKey MessageKey = "test_trigger_message_key"
	expectedResult := 5
	result := 0

	messageBus := NewBasicBus(&testErrorHandler{})
	messageBus.RegisterCallback(testDispatchMessageKey, func(c Message) error {
		result++
		return nil
	})
	messageBus.Listen()

	messageBus.Dispatch(Message{testDispatchMessageKey, MessageParams{}})
	messageBus.Dispatch(Message{testDispatchMessageKey, MessageParams{}})
	messageBus.Dispatch(Message{testDispatchMessageKey, MessageParams{}})
	messageBus.Dispatch(Message{testDispatchMessageKey, MessageParams{}})
	messageBus.Dispatch(Message{testDispatchMessageKey, MessageParams{}})
	time.Sleep(time.Millisecond * 5)

	if result != expectedResult {
		t.Errorf("Expected: '%d',  got: '%d'", expectedResult, result)
	}
}
