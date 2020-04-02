package account

import (
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"testing"
	"time"
)

func TestClient_Delete_Non_Existent_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(404, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("non-existent-id", 0)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "cannot find account non-existent-id")
}

func TestClient_Delete_Empty_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(204, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("", 0)

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Delete_Negative_Version(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(204, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("existing-id", -1)

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Delete_Wrong_Version(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(409, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("existing-id", 1)

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Delete_Server_Error(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(500, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("existing-id", 1)

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Delete_Valid_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(204, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	result, err := client.Delete("non-existent-id", 0)

	// Assert
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestClient_Delete_Wrong_Address(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://wrong-address",
		Client:    &http.Client{Timeout: time.Duration(1 * time.Second)},
		Validator: validator.New(),
	}

	// Act
	_, err := client.Delete("existing-id", 1)

	// Assert
	assert.NotNil(t, err)
}
