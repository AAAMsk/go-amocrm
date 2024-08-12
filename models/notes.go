package models

type Note struct {
	ID                int         `json:"id,omitempty"`
	ResponsibleUserID int         `json:"responsible_user_id,omitempty"`
	NoteParams        NoteParams  `json:"params,omitempty"`
	RequestID         interface{} `json:"request_id,omitempty"`
}

type NoteParams struct {
	Text     string `json:"text,omitempty"`
	NoteType string `json:"note_type,omitempty"`
}
