package teststore

import (
	"github.com/firkhraag/http-rest-api/internal/app/model"
	"github.com/firkhraag/http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	userRepository store.UserRepository
}

// NewStore ...
func NewStore() *Store {
	return &Store{}
}

// GetUserRepository ...
func (s *Store) GetUserRepository() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.userRepository
}
