package service

import "messageApi/internal/types"

// ServiceStub provides a stub for use in testing
type ServiceStub struct {
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
func (d *ServiceStub) CreateMessage(msg types.Message) (types.Message, error) {
	return d.CreateMessageResponse, d.CreateMessageError
}

// GetMessage returns static vars for use in testing
func (d *ServiceStub) GetMessage(id int) (types.Message, error) {
	return d.GetMessageResponse, d.GetMessageError
}

// ListMessages returns static vars for use in testing
func (d *ServiceStub) ListMessages() ([]types.Message, error) {
	return d.ListMessagesResponse, d.ListMessagesError
}

// UpdateMessage returns static vars for use in testing
func (d *ServiceStub) UpdateMessage(msg types.Message) (types.Message, error) {
	return d.UpdateMessageResponse, d.UpdateMessageError
}

// DeleteMessage returns static vars for use in testing
func (d *ServiceStub) DeleteMessage(id int) error {
	return d.DeleteMessageError
}
