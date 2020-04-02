package identification

type Private struct {
	Address        *string `json:"address"`
	BirthCountry   *string `json:"birth_country,omitempty"`
	BirthDate      *string `json:"birth_date,omitempty"`
	City           *string `json:"city,omitempty"`
	Country        *string `json:"country,omitempty"`
	DocumentNumber *string `json:"document_number,omitempty"`
	FirstName      *string `json:"first_name,omitempty"`
	LastName       *string `json:"last_name,omitempty"`
	Title          *string `json:"title,omitempty"`
}
