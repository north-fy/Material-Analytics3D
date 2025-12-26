package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate go test -coverprofile=coverage.out
//go:generate go tool cover -html=coverage.out -o coverage.html

func TestUser_AuthUser(t *testing.T) {
	tests := []struct {
		name    string
		model   User
		require error
	}{
		{
			name: "success",
			model: User{
				Login:    "login",
				Password: "password",
				Access:   AccessType{Access: AccessUser},
			},
			require: nil,
		},
		{
			name: "wrong login or password",
			model: User{
				Login:    "login2",
				Password: "password",
				Access:   AccessType{Access: AccessUser},
			},
			require: errWrongData,
		},
		{
			name: "wrong login or password",
			model: User{
				Login:    "login",
				Password: "password1",
				Access:   AccessType{Access: AccessUser},
			},
			require: errWrongData,
		},
	}

	modelUser := User{
		Login:    "login",
		Password: "password",
		Access:   AccessType{Access: AccessUser},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := modelUser.AuthUser(tc.model.Login, tc.model.Password)
			if err != nil {
				require.EqualError(t, err, tc.require.Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestUser_CheckAccessUser(t *testing.T) {
	tests := []struct {
		name    string
		model   User
		require bool
	}{
		{
			name: "success",
			model: User{
				Access: AccessType{Access: AccessExecutor},
			},
			require: true,
		},
		{
			name: "success",
			model: User{
				Access: AccessType{Access: AccessUnknown},
			},
			require: false,
		},
	}

	modelUser, err := NewUser(AccessType{Access: AccessExecutor}, "123", "123")
	assert.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flag := modelUser.CheckAccessUser(tc.model.Access)
			require.Equal(t, flag, tc.require)
		})
	}

}
