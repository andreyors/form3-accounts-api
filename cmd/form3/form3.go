package form3

import (
	Account "github.com/andreyors/accounts/cmd/form3/account"
	"time"
)

type Form3 struct {
	Account Account.ClientInterface
}

func New(host string, timeout time.Duration) *Form3 {
	return &Form3{
		Account: Account.NewClient(host, timeout),
	}
}
