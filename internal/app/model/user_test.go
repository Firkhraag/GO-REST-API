package model_test

import (
	"testing"

	"github.com/firkhraag/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				tu := model.TestUser(t)
				tu.Email = ""
				return tu
			},
			isValid: false,
		},
		{
			name: "wrong email",
			u: func() *model.User {
				tu := model.TestUser(t)
				tu.Email = "tralala"
				return tu
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				tu := model.TestUser(t)
				tu.Password = ""
				return tu
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				tu := model.TestUser(t)
				tu.Password = "123"
				return tu
			},
			isValid: false,
		},
		{
			name: "encrypted password",
			u: func() *model.User {
				tu := model.TestUser(t)
				tu.Password = ""
				tu.EncryptedPassword = "encrypted"
				return tu
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
