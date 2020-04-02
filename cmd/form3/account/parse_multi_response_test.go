package account

import (
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMultiResponse_Empty(t *testing.T) {
	collection, err := ParseMultiResponse(Test.MockResponse("empty.json"))

	assert.Nil(t, err)
	assert.Equal(t, 0, len(collection.Collection))
}

func TestParseMultiResponse_Wrong_One_Account(t *testing.T) {
	_, err := ParseMultiResponse(Test.MockResponse("one.json"))

	assert.NotNil(t, err)
}

func TestParseMultiResponse_Two_Accounts(t *testing.T) {
	collection, err := ParseMultiResponse(Test.MockResponse("two.json"))

	assert.Nil(t, err)
	assert.Equal(t, 2, len(collection.Collection))
	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", collection.Collection[0].ID.String())
	assert.Equal(t, "ea6239c1-99e9-42b3-bca1-92f5c068da6b", collection.Collection[1].ID.String())
}

func TestParseMultiResponse_Bad_Response(t *testing.T) {
	_, err := ParseMultiResponse(Test.BadResponse())

	assert.NotNil(t, err)
}
