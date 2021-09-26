package services

import (
	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/utils/cryptoUtils"
	"github.com/bookstore_users-api/utils/dateUtils"
	"github.com/bookstore_users-api/utils/errors"
)

var Userservice usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	GetUser(userId int64) (*users.User, *errors.RestErr)
	UpdateUSer(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
	FindByStatus(status string) (*[]users.User, *errors.RestErr)
}

func (us *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = "active"

	user.DateCreated = dateUtils.GetNowDb()
	user.Password = cryptoUtils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

func (us *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.StatusBadRequestError("invalid user id")
	}

	result := &users.User{
		Id: userId,
	}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}

func (us *usersService) UpdateUSer(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: user.Id,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			result.FirstName = user.FirstName
		}
		if user.LastName != "" {
			result.LastName = user.LastName
		}
		if user.Email != "" {
			result.Email = user.Email
		}

	} else {
		result.FirstName = user.FirstName
		result.LastName = user.LastName
		result.Email = user.Email

	}

	if err := result.Update(); err != nil {
		return nil, err
	}

	return result, nil

}

func (us *usersService) DeleteUser(userID int64) *errors.RestErr {

	result := &users.User{
		Id: userID,
	}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil
}

func (us *usersService) FindByStatus(status string) (*[]users.User, *errors.RestErr) {
	user := users.User{}
	result, err := user.FindUser(status)
	if err != nil {
		return nil, errors.StatusInternalServerError(err.Error)
	}
	return &result, nil
}
