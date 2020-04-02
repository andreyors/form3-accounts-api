package common

import (
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"net/url"
	"strconv"
)

func PageParams(pagination Model.Pagination) string {
	params := url.Values{}

	number := pagination.Number
	if number < 0 {
		number = 0
	}

	size := pagination.Size
	if size <= 0 {
		size = 100
	}

	params.Add("page[number]", strconv.Itoa(number))
	params.Add("page[size]", strconv.Itoa(size))

	return params.Encode()
}
