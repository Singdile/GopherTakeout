package repository

type OrderRepository interface {
	FindAll() ([]interface{}, error)
	FindByID(id int) (interface{}, error)
	FindByStatus(status int) ([]interface{}, error)
	Create(order interface{}) error
	UpdateStatus(id int, status int) error
}
