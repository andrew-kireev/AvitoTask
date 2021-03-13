package models

//easyjson:json
type Request struct {
	Date   string `json:"date"`
	Views  string `json:"views"`
	Clicks string `json:"clicks"`
	Cost   string `json:"cost"`
}
