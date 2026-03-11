package model

type Shop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Status      int    `json:"status"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
