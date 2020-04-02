package account

import (
	"github.com/andreyors/accounts/cmd/form3/models"
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"testing"
	"time"
)

func TestClient_List_Empty(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(200, Test.Sample("empty.json")),
		Validator: validator.New(),
	}

	// Act
	_, err := client.List(models.NewPagination(1, 10))

	// Assert
	assert.Nil(t, err)
}

func TestClient_List_Two_Accounts(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(200, Test.Sample("two.json")),
		Validator: validator.New(),
	}

	// Act
	result, err := client.List(models.NewPagination(1, 10))

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result.Collection))
}

func TestClient_List_Wrong_Server(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://wrong-address",
		Client:    &http.Client{Timeout: time.Duration(1 * time.Second)},
		Validator: validator.New(),
	}

	// Act
	_, err := client.List(models.NewPagination(1, 10))

	// Assert
	assert.NotNil(t, err)
}
