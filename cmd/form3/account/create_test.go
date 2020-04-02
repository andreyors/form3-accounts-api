package account

import (
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"github.com/andreyors/accounts/cmd/form3/models/account"
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"testing"
	"time"
)

func NewAccount() Model.Account {
	return Model.Account{
		ID:             uuid.New(),
		OrganisationID: uuid.New(),
		Type:           "accounts",
		Version:        0,
		Attributes: &account.Attributes{
			Country:    "GB",
			BankIDCode: "GBDSC",
			BIC:        "NWBKGB22",
			BankID:     "400300",
		},
	}
}

func TestClient_Create_Empty_Account(t *testing.T) {
	// Arrange
	client := NewClient("http://0.0.0.0:8080", time.Duration(5*time.Second))

	// Act
	_, err := client.Create(Model.Account{})

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Create_Wrong_Response_Code(t *testing.T) {
	// Arrange
	client := &Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(200, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	acct := NewAccount()
	_, err := client.Create(acct)

	// Assert
	assert.NotNil(t, err)
}

func TestClient_Create_Valid_Account(t *testing.T) {
	// Arrange
	client := Client{
		Host:      "http://0.0.0.0:8080",
		Client:    Test.MockClient(201, Test.Sample("one.json")),
		Validator: validator.New(),
	}

	// Act
	acct := NewAccount()
	account, err := client.Create(acct)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "accounts", account.Type)
	assert.Equal(t, 0, account.Version)
	assert.Equal(t, "GB", account.Attributes.Country)
}

func TestClient_Create_Wrong_Address(t *testing.T) {
	// Arrange
	client := Client{
		Host:      "http://wrong-address",
		Client:    &http.Client{Timeout: time.Duration(1 * time.Second)},
		Validator: validator.New(),
	}

	// Act
	acct := NewAccount()
	_, err := client.Create(acct)

	assert.NotNil(t, err)
}
