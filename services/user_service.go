package services

import (
	"github.com/marksfu1/bookstore_users-api/controllers/users/utils/errors"
	"github.com/marksfu1/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
