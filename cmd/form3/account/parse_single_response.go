package account

import (
	"encoding/json"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ParseSingleResponse(response *http.Response) (Model.Account, error) {
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return Model.Account{}, err
	}

	var serverResponse struct {
		Data  Model.Account `json:"data"`
		Links Model.Links   `json:"links"`
	}

	err = json.Unmarshal(content, &serverResponse)
	if err != nil {
		log.Error(err)
		return Model.Account{}, err
	}

	return serverResponse.Data, nil
}
