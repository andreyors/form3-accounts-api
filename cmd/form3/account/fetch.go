package account

import (
	"errors"
	"fmt"
	Common "github.com/andreyors/accounts/cmd/form3/common"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Client) Fetch(accountID string) (Model.Account, error) {
	var request *http.Request
	var response *http.Response
	var err error

	if accountID == "" {
		return Model.Account{}, errors.New("cannot process empty accountID")
	}

	uri := fmt.Sprintf("%s/%s/%s/%s", a.Host, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID)
	request, _ = http.NewRequest("GET", uri, nil)
	request.Header.Add("Accept", Common.JSON_TYPE)

	response, err = a.Client.Do(request)
	if err != nil {
		log.Error(err)
		return Model.Account{}, err
	}

	if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("wrong code %d", response.StatusCode)

		log.Error(errorMessage)
		return Model.Account{}, errors.New(errorMessage)
	}

	return ParseSingleResponse(response)
}
