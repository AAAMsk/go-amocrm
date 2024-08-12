package models

type Task struct {
	ID                int    `json:"id,omitempty"`
	EntityId          int    `json:"entity_id,omitempty"`
	TaskTypeId        int    `json:"task_type_id,omitempty"`
	ResponsibleUserID int    `json:"responsible_user_id,omitempty"`
	CompleteTill      int    `json:"complete_till,omitempty"`
	EntityType        string `json:"entity_type,omitempty"`
	Text              string `json:"text,omitempty"`
}
