package account

import (
	"errors"
	"fmt"
	Common "github.com/andreyors/accounts/cmd/form3/common"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Client) Delete(accountID string, version int) (bool, error) {
	var request *http.Request
	var response *http.Response
	var errorMessage string
	var err error

	if accountID == "" {
		return false, errors.New("cannot process empty accountID")
	}
	if version < 0 {
		return false, errors.New("cannot process negative version")
	}

	uri := fmt.Sprintf("%s/%s/%s/%s?version=%d", a.Host, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID, version)
	request, _ = http.NewRequest("DELETE", uri, nil)
	request.Header.Add("Accept", Common.JSON_TYPE)

	response, err = a.Client.Do(request)
	if err != nil {
		log.Error(err)
		return false, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNoContent:
		log.Debugf("account %s deleted", accountID)
		return true, nil
	case http.StatusNotFound:
		errorMessage = fmt.Sprintf("cannot find account %s", accountID)
	case http.StatusConflict:
		errorMessage = fmt.Sprintf("cannot process version %d for account %s", version, accountID)
	default:
		errorMessage = fmt.Sprintf("unknown status code %d", response.StatusCode)
	}

	return false, errors.New(errorMessage)
}
