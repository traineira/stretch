package model

type Story struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	Category string `json:"text"`
	Finished bool   `json:"done"`
	UserID   string `json:"user"`
}
