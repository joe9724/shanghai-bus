package models

type UpDown struct {
	id int64 `json:"id"`
	LineID string `json:"line_id,omitempty"`
	UpDown int64 `json:"updown"`
	StartStation string `json:"start_station"`
	EndStation string `json:"end_station"`
}

