package identification

type Representative struct {
	BirthDate *string `json:"birth_date,omitempty"`
	Name      *string `json:"name,omitempty"`
	Residency *string `json:"residency,omitempty"`
}
