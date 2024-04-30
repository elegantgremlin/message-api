package database

import "messageApi/internal/types"

// DatabaseStub provides a stub for use in testing
type DatabaseStub struct {
	GetMessageResponse    types.Message
	GetMessageError       error
	ListMessagesResponse  []types.Message
	ListMessagesError     error
	CreateMessageResponse types.Message
	CreateMessageError    error
	UpdateMessageResponse types.Message
	UpdateMessageError    error
	DeleteMessageError    error
}

// CreateMessage returns static vars for use in testing
func (d *DatabaseStub) CreateMessage(msg types.Message) (types.Message, error) {
	return d.CreateMessageResponse, d.CreateMessageError
}

// GetMessage returns static vars for use in testing
func (d *DatabaseStub) GetMessage(id int) (types.Message, error) {
	return d.GetMessageResponse, d.GetMessageError
}

// ListMessages returns static vars for use in testing
func (d *DatabaseStub) ListMessages() ([]types.Message, error) {
	return d.ListMessagesResponse, d.ListMessagesError
}

// UpdateMessage returns static vars for use in testing
func (d *DatabaseStub) UpdateMessage(msg types.Message) (types.Message, error) {
	return d.UpdateMessageResponse, d.UpdateMessageError
}

// DeleteMessage returns static vars for use in testing
func (d *DatabaseStub) DeleteMessage(id int) error {
	return d.DeleteMessageError
}
