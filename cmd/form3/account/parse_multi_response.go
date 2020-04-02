package account

import (
	"encoding/json"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ParseMultiResponse(response *http.Response) (Model.Accounts, error) {
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return Model.Accounts{}, err
	}

	var serverResponse struct {
		Data  []Model.Account `json:"data"`
		Links Model.Links     `json:"links"`
	}

	err = json.Unmarshal(content, &serverResponse)
	if err != nil {
		log.Error(err)
		return Model.Accounts{}, err
	}

	collection := Model.Accounts{}
	for _, acct := range serverResponse.Data {
		collection.Collection = append(collection.Collection, acct)
	}

	return collection, nil
}
