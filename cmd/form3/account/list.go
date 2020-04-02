package account

import (
	"fmt"
	Common "github.com/andreyors/accounts/cmd/form3/common"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func (a *Client) List(pagination Model.Pagination) (Model.Accounts, error) {
	var request *http.Request
	var response *http.Response
	var uri *url.URL
	var err error

	targetUri := fmt.Sprintf("%s/%s/%s", a.Host, Common.API_VERSION, Common.ACCOUNT_ROUTE)
	uri, _ = url.Parse(targetUri)
	uri.RawQuery = Common.PageParams(pagination)

	request, _ = http.NewRequest("GET", uri.String(), nil)
	request.Header.Add("Accept", Common.JSON_TYPE)

	response, err = a.Client.Do(request)
	if err != nil {
		log.Error(err)
		return Model.Accounts{}, err
	}

	return ParseMultiResponse(response)
}
