package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	Common "github.com/andreyors/accounts/cmd/form3/common"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"github.com/andreyors/accounts/cmd/form3/validation"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Client) Create(account Model.Account) (Model.Account, error) {
	var request *http.Request
	var response *http.Response
	var payload []byte
	var err error

	a.Validator.RegisterStructValidation(validation.GBBankIDCodeValidation, Model.Account{})

	err = a.Validator.Struct(account)
	if err != nil {
		log.Error(err)

		return Model.Account{}, err
	}

	type createAccount struct {
		Data Model.Account `json:"data"`
	}
	data := &createAccount{Data: account}

	payload, _ = json.Marshal(data)

	uri := fmt.Sprintf("%s/%s/%s", a.Host, Common.API_VERSION, Common.ACCOUNT_ROUTE)
	request, _ = http.NewRequest("POST", uri, bytes.NewBuffer(payload))
	request.Header.Add("Content-Type", Common.JSON_TYPE)

	response, err = a.Client.Do(request)
	if err != nil {
		log.Error(err)
		return Model.Account{}, err
	}

	if response.StatusCode != http.StatusCreated {
		errorMessage := fmt.Sprintf("wrong code %d", response.StatusCode)

		log.Error(errorMessage)
		return Model.Account{}, errors.New(errorMessage)
	}

	return ParseSingleResponse(response)
}
