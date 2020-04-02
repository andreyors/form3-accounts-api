package models

import (
	Model "github.com/andreyors/accounts/cmd/form3/models/account"
	"github.com/google/uuid"
)

type Account struct {
	Attributes     *Model.Attributes    `json:"attributes,omitempty"`
	ID             uuid.UUID            `json:"id" validate:"required"`
	OrganisationID uuid.UUID            `json:"organisation_id" validate:"required"`
	Relationships  *Model.Relationships `json:"relationships,omitempty"`
	Type           string               `json:"type"`
	Version        int                  `json:"version"`
}

func NewAccount() *Account {
	return &Account{}
}
