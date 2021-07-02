package model

type Restaurant struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Phone string `json:"phone"`
}
