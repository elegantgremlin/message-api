package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestErrorAsJSON tests converting a string into a map successfully
func TestErrorAsJSON(t *testing.T) {
	msg := "an error has occurred"
	msgJSON := map[string]string{"error": msg}

	jsonError := errorAsJSON(msg)

	assert.Equal(t, msgJSON, jsonError)
}
