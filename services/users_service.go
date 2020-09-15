package services

import (
	"github.com/m90trash/bookstore_users-api/domain/users"
	"github.com/m90trash/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {

	return &user, nil
}
