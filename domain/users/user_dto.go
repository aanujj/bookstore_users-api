package users

import (
	"fmt"
	"strings"

	"github.com/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.StatusNotFoundError(fmt.Sprintf("invalid email %v address", user.Email))
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.StatusNotFoundError(fmt.Sprintf("invalid password %v address", user.Password))
	}
	return nil
}
