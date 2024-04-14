package models

type Data struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Company     string  `json:"company"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}
