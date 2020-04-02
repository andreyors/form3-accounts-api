package validation

import (
	Model "github.com/andreyors/accounts/cmd/form3/models"
	"gopkg.in/go-playground/validator.v9"
)

func GBBankIDCodeValidation(sl validator.StructLevel) {
	account := sl.Current().Interface().(Model.Account)

	if account.Attributes != nil && account.Attributes.Country == "GB" {
		if account.Attributes.BankIDCode != "GBDSC" {
			sl.ReportError(account.Attributes.BankIDCode, "bank_id_code", "BankIDCode", "gbbankidcode", "")
		}

		if account.Attributes.BIC == "" {
			sl.ReportError(account.Attributes.BIC, "bic", "BIC", "gbbic", "")
		}

		if account.Attributes.BankID == "" {
			sl.ReportError(account.Attributes.BankID, "bank_id", "BankID", "gbbankid", "")
		}
	}
}
