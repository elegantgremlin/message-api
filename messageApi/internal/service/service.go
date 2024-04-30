package service

import (
	"errors"
	"messageApi/internal/database"
	"messageApi/internal/types"
)

// Service is the interface for the service module of the application
type Service interface {
	CreateMessage(types.Message) (types.Message, error)
	ListMessages() ([]types.Message, error)
	GetMessage(int) (types.Message, error)
	UpdateMessage(types.Message) (types.Message, error)
	DeleteMessage(int) error
}

// service is the implementation of the service module
type service struct {
	Db database.Database
}

// NewService creates an instance of the service module
func NewService(cfg types.Config, db database.Database) (Service, error) {
	return &service{db}, nil
}

// CreateMessage validates the message then sends it to the data module
func (s *service) CreateMessage(msg types.Message) (types.Message, error) {
	if err := validateMessage(msg); err != nil {
		return types.Message{}, err
	}

	msg.IsPalindrome = isPalindrome(msg.Message)

	msg, err := s.Db.CreateMessage(msg)
	if err != nil {
		return types.Message{}, err
	}

	return msg, nil
}

// isPalindrome checks if the message is a palindrome. Needs to be an exact match, including casing and whitespace
func isPalindrome(msg string) bool {
	i := 0
	j := len(msg) - 1

	for (i <= j) && msg[i] == msg[j] {
		i++
		j--
	}

	return (i > j)
}

// validateMessage checks that the Message matches the requirements
func validateMessage(msg types.Message) error {
	if len(msg.Message) > 100 {
		return errors.New("message cannot be longer than 100 characters")
	}

	if len(msg.Message) == 0 {
		return errors.New("message cannot be an empty string")
	}

	return nil
}

// GetMessage returns a single Message from the data module
func (s *service) GetMessage(id int) (types.Message, error) {
	return s.Db.GetMessage(id)
}

// ListMessages returns a list of Messages from the data module
func (s *service) ListMessages() ([]types.Message, error) {
	return s.Db.ListMessages()
}

// UpdateMessage updates an existing Message
func (s *service) UpdateMessage(msg types.Message) (types.Message, error) {
	if err := validateMessage(msg); err != nil {
		return types.Message{}, err
	}

	msg.IsPalindrome = isPalindrome(msg.Message)

	msg, err := s.Db.UpdateMessage(msg)
	if err != nil {
		return types.Message{}, err
	}

	return msg, nil
}

// DeleteMessage deletes an existing message
func (s *service) DeleteMessage(id int) error {
	return s.Db.DeleteMessage(id)
}
