package models

//easyjson:json
type Request struct {
	Date   string `json:"date" validate:"required"`
	Views  string `json:"views" validate:"required"`
	Clicks string `json:"clicks" validate:"required"`
	Cost   string `json:"cost" validate:"required"`
}
