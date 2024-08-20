package models

type CustomsFields struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Code  string  `json:"code,omitempty"`
	Enums []Enums `json:"enums,omitempty"`
}

type Enums struct {
	ID    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Sort  int    `json:"sort,omitempty"`
}
