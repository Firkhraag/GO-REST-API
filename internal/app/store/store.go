package store

// Store ...
type Store interface {
	GetUserRepository() UserRepository
}
