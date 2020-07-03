package teststore

import (
	"github.com/firkhraag/http-rest-api/internal/app/model"
	"github.com/firkhraag/http-rest-api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// CreateUser ...
func (r *UserRepository) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// FindUserByID ...
func (r *UserRepository) FindUserByID(id int) (*model.User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, store.ErrRecordNotFound
}

// FindUserByEmail ...
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
