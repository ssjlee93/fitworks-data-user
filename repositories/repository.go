package repositories

import "github.com/ssjlee93/fitworks-data-user/dtos"

type Repository[T dtos.Role | dtos.User] interface {
	// Create queries DB to create a row
	// return the created obj
	Create(t T) (*T, error)
	// ReadOne queries DB to retrieve a row by id
	// returns the retrieved obj
	ReadOne(id int64) (*T, error)
	// ReadAll queries DB to retrieve all rows in the table
	// returns a slice of objs
	ReadAll() ([]T, error)
	// Update queries DB to update a row
	// returns the updated obj
	Update(t T) (*T, error)
	// Delete queries DB to hard delete a row
	// returns the deleted obj
	Delete(id int64) (*T, error)
}
