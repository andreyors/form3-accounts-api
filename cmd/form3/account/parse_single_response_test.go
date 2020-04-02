package account

import (
	Test "github.com/andreyors/accounts/cmd/form3/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSingleResponse_Empty(t *testing.T) {
	_, err := ParseSingleResponse(Test.MockResponse("empty.json"))

	assert.NotNil(t, err)
}

func TestParseSingleResponse_One_Account(t *testing.T) {
	account, err := ParseSingleResponse(Test.MockResponse("one.json"))

	assert.Nil(t, err)
	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", account.ID.String())
}

func TestParseSingleResponse_Wrong_Two_Accounts(t *testing.T) {
	_, err := ParseSingleResponse(Test.MockResponse("two.json"))

	assert.NotNil(t, err)
}

func TestParseSingleResponse_Bad_Response(t *testing.T) {
	_, err := ParseSingleResponse(Test.BadResponse())

	assert.NotNil(t, err)
}
