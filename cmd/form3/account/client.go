package account

import (
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type Client struct {
	Host      string
	Client    *http.Client
	Validator *validator.Validate
}

type ClientInterface interface {
	Create(account Model.Account) (Model.Account, error)
	Delete(accountID string, version int) (bool, error)
	Fetch(accountID string) (Model.Account, error)
	List(pagination Model.Pagination) (Model.Accounts, error)
}

func NewClient(host string, timeout time.Duration) *Client {
	return &Client{
		Host:      host,
		Client:    &http.Client{Timeout: timeout},
		Validator: validator.New(),
	}
}
