package sqlstore

import (
	"database/sql"

	"github.com/firkhraag/http-rest-api/internal/app/model"
	"github.com/firkhraag/http-rest-api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// CreateUser ...
func (r *UserRepository) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindUserByID ...
func (r *UserRepository) FindUserByID(id int) (*model.User, error) {
	u := &model.User{
		ID: id,
	}
	err := r.store.db.QueryRow(
		"SELECT email, encrypted_password FROM users WHERE id = $1", id,
	).Scan(&u.Email, &u.EncryptedPassword)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}
	return u, err
}

// FindUserByEmail ...
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	u := &model.User{
		Email: email,
	}
	err := r.store.db.QueryRow(
		"SELECT id, encrypted_password FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.EncryptedPassword)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}
	return u, err
}
