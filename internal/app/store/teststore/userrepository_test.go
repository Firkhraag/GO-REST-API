package teststore_test

import (
	"testing"

	"github.com/firkhraag/http-rest-api/internal/app/model"
	"github.com/firkhraag/http-rest-api/internal/app/store"
	"github.com/firkhraag/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUserRepository(t *testing.T) {
	assert.NoError(t, teststore.NewStore().GetUserRepository().CreateUser(model.TestUser(t)))
}

func TestUserRepository_FindUserByID(t *testing.T) {
	s := teststore.NewStore()

	u := model.TestUser(t)
	_, err := s.GetUserRepository().FindUserByID(u.ID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	assert.NoError(t, s.GetUserRepository().CreateUser(u))

	u2, err := s.GetUserRepository().FindUserByID(u.ID)

	assert.NoError(t, err)
	assert.Equal(t, u2.ID, u.ID)
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	s := teststore.NewStore()

	u := model.TestUser(t)
	_, err := s.GetUserRepository().FindUserByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	assert.NoError(t, s.GetUserRepository().CreateUser(u))

	u2, err := s.GetUserRepository().FindUserByEmail(u.Email)

	assert.NoError(t, err)
	assert.Equal(t, u2.Email, u.Email)
}
