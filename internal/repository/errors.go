package repository

import (
	"fmt"
)

type NotFoundByIDError struct {
	ID       int
	Resource string
}

func (e *NotFoundByIDError) Error() string {
	return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
}
