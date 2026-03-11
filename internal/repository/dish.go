package repository

type DishRepository interface {
	FindAll() ([]interface{}, error)
	FindByID(id int) (interface{}, error)
	FindByCategoryID(categoryID int) ([]interface{}, error)
	Create(dish interface{}) error
	Update(dish interface{}) error
	Delete(id int) error
}
