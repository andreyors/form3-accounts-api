package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	// Arrange
	// Act
	client := NewClient("http://0.0.0.0:8080", time.Duration(5*time.Second))

	// Assert
	assert.IsType(t, &Client{}, client)
}
