package account

import (
	Relationship "github.com/andreyors/accounts/cmd/form3/models/relationship"
)

type Relationships struct {
	AccountEvents *Relationship.Links `json:"account_events,omitempty"`
	MasterAccount *Relationship.Links `json:"master_account,omitempty"`
}
