package identification

type Organisation struct {
	Address            *string         `json:"address"`
	City               *string         `json:"city,omitempty"`
	Country            *string         `json:"country,omitempty"`
	Name               *string         `json:"name,omitempty"`
	RegistrationNumber *string         `json:"registration_number,omitempty"`
	Representative     *Representative `json:"representative"`
	TaxResidency       *string         `json:"tax_residency,omitempty"`
}
