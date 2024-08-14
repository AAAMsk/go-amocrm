package models

type Contact struct {
	ID                 int              `json:"id,omitempty"`
	IsMain             bool             `json:"is_main,omitempty"`
	Name               string           `json:"name,omitempty"`
	FirstName          string           `json:"first_name,omitempty"`
	LastName           string           `json:"last_name,omitempty"`
	GroupID            int              `json:"group_id,omitempty"`
	ResponsibleUserID  int              `json:"responsible_user_id,omitempty"`
	CreatedBy          int              `json:"created_by,omitempty"`
	UpdatedBy          int              `json:"updated_by,omitempty"`
	CreatedAt          int              `json:"created_at,omitempty"`
	UpdatedAt          int              `json:"updated_at,omitempty"`
	ClosestTaskAt      interface{}      `json:"closest_task_at,omitempty"`
	IsDeleted          bool             `json:"is_deleted,omitempty"`
	IsUnsorted         bool             `json:"is_unsorted,omitempty"`
	AccountID          int              `json:"account_id,omitempty"`
	CustomFieldsValues []CustomFields   `json:"custom_fields_values,omitempty"`
	Embedded           *ContactEmbedded `json:"_embedded,omitempty"`
	Links              *LinkResponse    `json:"_links,omitempty"`
}

type SetMain struct {
	MainID   int `json:"main_id,omitempty"`
	LinkedID int `json:"linked_id,omitempty"`
}

//https://wachecking.amocrm.ru/ajax/linked/leads/set/main/contacts
//application/x-www-form-urlencoded

type ContactEmbedded struct {
	Tags            []Tag             `json:"tags,omitempty"`
	Leads           []Lead            `json:"leads,omitempty"`
	Companies       []Company         `json:"companies,omitempty"`
	CatalogElements []CatalogElements `json:"catalog_elements,omitempty"`
}
