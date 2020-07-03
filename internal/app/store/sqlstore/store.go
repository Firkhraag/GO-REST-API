package sqlstore

import (
	"database/sql"

	"github.com/firkhraag/http-rest-api/internal/app/store"
	_ "github.com/lib/pq" // Postgresql driver
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository store.UserRepository
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserRepository ...
func (s *Store) GetUserRepository() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
