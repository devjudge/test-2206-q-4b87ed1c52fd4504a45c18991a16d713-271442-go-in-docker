package model

type Lead struct {
	Id              int    `json:"id",omitempty`
	First_name      string `json:"first_name",omitempty`
	Last_name       string `json:"last_name",omitempty`
	Mobile          string `json:"mobile",omitempty`
	Email           string `json:"email",omitempty`
	Location_type   string `json:"location_type",omitempty`
	Location_string string `json:"location_string",omitempty`
	Status          string `json:"status",omitempty`
}
