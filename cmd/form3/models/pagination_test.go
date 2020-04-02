package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPagination(t *testing.T) {
	pager := NewPagination(1, 10)

	assert.IsType(t, Pagination{}, pager)
}
