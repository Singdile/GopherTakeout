package repository

type CategoryRepository interface {
	FindAll() ([]interface{}, error)
	FindByID(id int) (interface{}, error)
	Create(category interface{}) error
	Update(category interface{}) error
	Delete(id int) error
}
