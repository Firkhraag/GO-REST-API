package store

import "github.com/firkhraag/http-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	CreateUser(*model.User) error
	FindUserByID(int) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
}
