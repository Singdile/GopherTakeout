package repository

type ShopRepository interface {
	FindByID(id int) (interface{}, error)
	Update(shop interface{}) error
	UpdateStatus(id int, status int) error
}
