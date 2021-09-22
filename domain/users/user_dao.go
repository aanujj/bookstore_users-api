package users

import (
	"fmt"

	"github.com/bookstore_users-api/utils/errors"
)

//this function is going to save user in the database

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDb[user.Id]
	if result == nil {
		return errors.StatusNotFoundError(fmt.Sprintf("user id %v not found", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	return nil

}

func (user *User) Save() *errors.RestErr {

	if userDb[user.Id] != nil {
		if userDb[user.Id].Email == user.Email {
			errors.StatusBadRequestError(fmt.Sprintf("Email %v already exists", user.Email))
		}
		return errors.StatusBadRequestError(fmt.Sprintf("user %v already exists", user.Id))
	}
	userDb[user.Id] = user
	return nil
}
