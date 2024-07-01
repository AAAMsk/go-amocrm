package models

type PipelineStatus struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PipelineId int    `json:"pipeline_id,omitempty"`
}
