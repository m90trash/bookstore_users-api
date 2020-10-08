package users

import "github.com/m90trash/bookstore_users-api/utils/errors"

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	CreateDate string `json:"create_date"`
}

func (u *User) Validate() *errors.RestError {
	if u.Email == "" {
		return errors.NewBadRequestError("bad email")
	}

	return nil
}
