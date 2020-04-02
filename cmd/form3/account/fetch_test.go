package account

import (
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"testing"
	"time"
)

func TestClient_Fetch_Non_Existent_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(404, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Fetch("non-existent-id")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "wrong code 404")
}

func TestClient_Fetch_Empty_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(404, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.Fetch("")

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Fetch_Valid_Account(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(200, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	result, err := client.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", result.ID.String())
}

func TestClient_Fetch_Wrong_Server(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://wrong-address",
		Client:    &http.Client{Timeout: time.Duration(1 * time.Second)},
		Validator: validator.New(),
	}

	// Act
	_, err := client.Fetch("existing-id")

	// Assert
	assert.NotNil(t, err)
}
