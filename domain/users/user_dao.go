package users

import (
	"fmt"
	"github.com/m90trash/bookstore_users-api/utils/errors"
)

var repo = make(map[int64]*User)

func (u *User) Get() *errors.RestError {
	ur := repo[u.Id]

	if ur == nil {
		return errors.NewNotFoundError(fmt.Sprintf("id %d not found", u.Id))
	}

	u.Id = ur.Id
	u.Email = ur.Email
	u.CreateDate = ur.CreateDate
	u.FirstName = ur.FirstName
	u.LastName = ur.LastName

	return nil
}

func (u *User) Save() *errors.RestError {

	if repo[u.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("id %d jau yra", u.Id))
	}

	repo[u.Id] = u

	return nil

}
