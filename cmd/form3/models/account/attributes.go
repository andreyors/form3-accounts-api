package account

import (
	Identification "github.com/andreyors/accounts/cmd/form3/models/identification"
)

type Attributes struct {
	AccountClassification       *string                      `json:"account_classification,omitempty"`
	AccountMatchingOptOut       *bool                        `json:"account_matching_opt_out,omitempty"`
	AccountNumber               *string                      `json:"account_number,omitempty"`
	AlternativeBankAccountNames []string                     `json:"alternative_bank_account_names,omitempty"`
	BankAccountName             *string                      `json:"bank_account_name,omitempty"`
	BankID                      string                       `json:"bank_id,omitempty"`
	BankIDCode                  string                       `json:"bank_id_code,omitempty" validate:"required_with=country"`
	BaseCurrency                *string                      `json:"base_currency,omitempty"`
	BIC                         string                       `json:"bic,omitempty"`
	CustomerID                  *string                      `json:"customer_id,omitempty"`
	FirstName                   *string                      `json:"first_name,omitempty"`
	IBAN                        *string                      `json:"iban,omitempty"`
	JointAccount                *bool                        `json:"joint_account,omitempty"`
	OrganisationIdentification  *Identification.Organisation `json:"organisation_identification,omitempty"`
	PrivateIdentification       *Identification.Private      `json:"private_identification,omitempty"`
	SecondaryIdentification     *string                      `json:"secondary_identification,omitempty"`
	Status                      *string                      `json:"status,omitempty"`
	Switched                    *bool                        `json:"switched,omitempty"`
	Title                       *string                      `json:"title,omitempty"`
	Country                     string                       `json:"country,omitempty" validate:"required"`
}
