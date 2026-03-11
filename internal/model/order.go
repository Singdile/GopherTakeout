package model

type Order struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shop_id"`
	OrderNumber string  `json:"order_number"`
	Status      int     `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
