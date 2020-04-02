// +build e2e

package form3

import (
	Form3 "github.com/andreyors/accounts/cmd/form3"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"github.com/andreyors/accounts/cmd/form3/models/account"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_E2E_Create_User(t *testing.T) {
	api := os.Getenv("FORM3_API")

	if api == "" {
		t.Skip("Cannot test due to empty FORM3_API")
	}
	f3 := Form3.New(api, time.Duration(1*time.Second))

	accountID := uuid.New()
	acct := Model.Account{
		ID:             accountID,
		OrganisationID: uuid.New(),
		Type:           "accounts",
		Version:        0,
		Attributes: &account.Attributes{
			Country: "GB",
		},
	}

	result, err := f3.Account.Create(acct)

	assert.Nil(t, err)
	assert.Equal(t, accountID, result.ID)
}

func Test_E2E_Fetch_Fresh_User(t *testing.T) {
	api := os.Getenv("FORM3_API")
	if api == "" {
		t.Skip("Cannot test due to empty FORM3_API")
	}
	f3 := Form3.New(api, time.Duration(1*time.Second))

	accountID := uuid.New()
	acct := Model.Account{
		ID:             accountID,
		OrganisationID: uuid.New(),
		Type:           "accounts",
		Version:        0,
		Attributes: &account.Attributes{
			Country: "GB",
		},
	}

	f3.Account.Create(acct)
	time.Sleep(1 * time.Second)
	result, err := f3.Account.Fetch(accountID.String())

	assert.Nil(t, err)
	assert.Equal(t, accountID, result.ID)
}

func Test_E2E_List_Fresh_User(t *testing.T) {
	api := os.Getenv("FORM3_API")
	if api == "" {
		t.Skip("Cannot test due to empty FORM3_API")
	}
	f3 := Form3.New(api, time.Duration(1*time.Second))

	accountID := uuid.New()
	acct := Model.Account{
		ID:             accountID,
		OrganisationID: uuid.New(),
		Type:           "accounts",
		Version:        0,
		Attributes: &account.Attributes{
			Country: "GB",
		},
	}

	f3.Account.Create(acct)
	time.Sleep(1 * time.Second)
	result, err := f3.Account.List(Model.NewPagination(0, 1000))

	found := false
	if len(result.Collection) > 0 {
		for _, ac := range result.Collection {
			if ac.ID.String() == accountID.String() {
				found = true
				break
			}
		}
	}

	assert.Nil(t, err)
	assert.Less(t, 0, len(result.Collection))
	assert.True(t, found)
}

func Test_E2E_Delete_Fresh_User(t *testing.T) {
	api := os.Getenv("FORM3_API")
	if api == "" {
		t.Skip("Cannot test due to empty FORM3_API")
	}
	f3 := Form3.New(api, time.Duration(1*time.Second))

	accountID := uuid.New()
	acct := Model.Account{
		ID:             accountID,
		OrganisationID: uuid.New(),
		Type:           "accounts",
		Version:        0,
		Attributes: &account.Attributes{
			Country: "GB",
		},
	}

	f3.Account.Create(acct)
	time.Sleep(1 * time.Second)
	result, err := f3.Account.Delete(accountID.String(), 0)

	assert.Nil(t, err)
	assert.True(t, result)
}
