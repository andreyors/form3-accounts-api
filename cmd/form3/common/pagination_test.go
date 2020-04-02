package common

import (
	"fmt"
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPageParams_Empty_Pager(t *testing.T) {
	pager := Model.Pagination{}

	result := PageParams(pager)

	fmt.Println(result)

	assert.Equal(t, "page%5Bnumber%5D=0&page%5Bsize%5D=100", result)
}

func TestPageParams_Wrong_Page_Number(t *testing.T) {
	pager := Model.Pagination{
		Number: -1,
		Size:   10,
	}

	result := PageParams(pager)

	fmt.Println(result)

	assert.Equal(t, "page%5Bnumber%5D=0&page%5Bsize%5D=10", result)
}

func TestPageParams_Wrong_Page_Size(t *testing.T) {
	pager := Model.Pagination{
		Number: -1,
		Size:   -1,
	}

	result := PageParams(pager)

	fmt.Println(result)

	assert.Equal(t, "page%5Bnumber%5D=0&page%5Bsize%5D=100", result)
}

func TestPageParams_Valid_Page(t *testing.T) {
	pager := Model.Pagination{
		Number: 1,
		Size:   2,
	}

	result := PageParams(pager)

	fmt.Println(result)

	assert.Equal(t, "page%5Bnumber%5D=1&page%5Bsize%5D=2", result)
}
