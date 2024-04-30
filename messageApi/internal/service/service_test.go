package service

import (
	"errors"
	"messageApi/internal/database"
	"messageApi/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateMessagePalindrome tests creating a Message that is a palindrome
func TestCreateMessagePalindrome(t *testing.T) {
	input_msg := types.Message{Id: 0, Message: "racecar", IsPalindrome: false}
	output_msg := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	db_stub := database.DatabaseStub{
		CreateMessageResponse: output_msg,
		CreateMessageError:    nil,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.CreateMessage(input_msg)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, nil, err)
}

// TestCreateMessageNotPalindrome tests creating a Message that is not a palindrome
func TestCreateMessageNotPalindrome(t *testing.T) {
	input_msg := types.Message{Id: 0, Message: "Test", IsPalindrome: false}
	output_msg := types.Message{Id: 1, Message: "Test", IsPalindrome: false}
	db_stub := database.DatabaseStub{
		CreateMessageResponse: output_msg,
		CreateMessageError:    nil,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.CreateMessage(input_msg)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, nil, err)
}

// TestCreateMessageError tests returning an error when the call to the data module fails
func TestCreateMessageError(t *testing.T) {
	input_msg := types.Message{Id: 0, Message: "Test", IsPalindrome: false}
	output_err := errors.New("error creating new message")

	db_stub := database.DatabaseStub{
		CreateMessageResponse: types.Message{},
		CreateMessageError:    output_err,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.CreateMessage(input_msg)

	assert.Equal(t, types.Message{}, msg)
	assert.Equal(t, output_err, err)
}

// TestIsPalindrome tests whether various strings are palindromes
func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, false, isPalindrome("test"))
	assert.Equal(t, true, isPalindrome("racecar"))
	assert.Equal(t, false, isPalindrome("race car"))
	assert.Equal(t, false, isPalindrome("Racecar"))
}

// TestUpdateMessagePalindrome tests updating a Message that is a palindrome
func TestUpdateMessagePalindrome(t *testing.T) {
	input_msg := types.Message{Id: 1, Message: "racecar", IsPalindrome: false}
	output_msg := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	db_stub := database.DatabaseStub{
		UpdateMessageResponse: output_msg,
		UpdateMessageError:    nil,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.UpdateMessage(input_msg)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, nil, err)
}

// TestUpdateMessageNotPalindrome tests updating a Message that is not a palindrome
func TestUpdateMessageNotPalindrome(t *testing.T) {
	input_msg := types.Message{Id: 1, Message: "Test", IsPalindrome: false}
	output_msg := types.Message{Id: 1, Message: "Test", IsPalindrome: false}
	db_stub := database.DatabaseStub{
		UpdateMessageResponse: output_msg,
		UpdateMessageError:    nil,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.UpdateMessage(input_msg)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, nil, err)
}

// TestUpdateMessageError tests error is returned when the call to the data module fails
func TestUpdateMessageError(t *testing.T) {
	input_msg := types.Message{Id: 1, Message: "Test", IsPalindrome: false}
	output_err := errors.New("error updating message")

	db_stub := database.DatabaseStub{
		UpdateMessageResponse: types.Message{},
		UpdateMessageError:    output_err,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.UpdateMessage(input_msg)

	assert.Equal(t, types.Message{}, msg)
	assert.Equal(t, output_err, err)
}

// TestGetMessage tests retrieving a single Message returns values from data module
func TestGetMessage(t *testing.T) {
	output_msg := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	output_err := errors.New("error updating message")

	db_stub := database.DatabaseStub{
		GetMessageResponse: output_msg,
		GetMessageError:    output_err,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.GetMessage(output_msg.Id)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, output_err, err)
}

// TestListMessage tests retrieving a list of Messages returns values from data module
func TestListMessage(t *testing.T) {
	output_msg := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	output_err := errors.New("error updating message")

	db_stub := database.DatabaseStub{
		GetMessageResponse: output_msg,
		GetMessageError:    output_err,
	}
	service, _ := NewService(types.Config{}, &db_stub)

	msg, err := service.GetMessage(output_msg.Id)

	assert.Equal(t, output_msg, msg)
	assert.Equal(t, output_err, err)
}

// TestDeleteMessage tests deleting a single Message returns value from data module
func TestDeleteMessage(t *testing.T) {
	id := 1
	output_err := errors.New("error updating message")

	db_stub := database.DatabaseStub{DeleteMessageError: output_err}
	service, _ := NewService(types.Config{}, &db_stub)

	err := service.DeleteMessage(id)

	assert.Equal(t, output_err, err)
}

// TestValidateMessageLength tests validation fails if Message is longer than 100 characters
func TestValidateMessageLength(t *testing.T) {
	output_err := errors.New("message cannot be longer than 100 characters")
	msg_text := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	msg := types.Message{Id: 0, Message: msg_text, IsPalindrome: true}

	err := validateMessage(msg)

	assert.Equal(t, output_err, err)
}

// TestValidateMessageEmptyMessage tests validation fails if Message is longer than 100 characters
func TestValidateMessageEmptyMessage(t *testing.T) {
	output_err := errors.New("message cannot be an empty string")
	msg := types.Message{Id: 0, Message: "", IsPalindrome: true}

	err := validateMessage(msg)

	assert.Equal(t, output_err, err)
}
