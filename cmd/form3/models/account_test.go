package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount()

	assert.IsType(t, &Account{}, account)
}
