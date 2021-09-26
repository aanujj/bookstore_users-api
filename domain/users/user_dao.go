package users

import (
	"strings"

	usersdb "github.com/bookstore_users-api/datasources/mysql/users_db"
	"github.com/bookstore_users-api/utils/dateUtils"
	"github.com/bookstore_users-api/utils/errors"
)

//this function is going to save user in the database

const (
	queryInsertUser       = "insert into users(first_name,last_name,email,date_created,status,password) values (?,?,?,?,?,?);"
	queryGetUser          = "select id,first_name,last_name,email,date_created,status from users where id=?;"
	queryUpdateUser       = "update users set first_name=?,last_name=?,email=? where id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "select id,first_name,last_name,email,date_created,status from users where status=?;"
)

func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	user.DateCreated = dateUtils.GetNowDb()
	GetResult := usersdb.Client.QueryRow(queryGetUser, user.Id)
	if err := GetResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	return nil

}

func (user *User) Save() *errors.RestErr {

	// stmt, err := usersdb.Client.Prepare(queryInsertUser)
	// if err != nil {
	// 	return errors.StatusInternalServerError(err.Error())
	// }

	//InsertResult2 is same as above with short form
	user.DateCreated = dateUtils.GetNowDb()
	InsertResult2, err := usersdb.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	//InsertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "users.email") {
			return errors.StatusBadRequestError("email already exists")

		}
		return errors.StatusInternalServerError(err.Error())
	}
	userId, err := InsertResult2.LastInsertId()

	if err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	user.Id = userId
	// if userDb[user.Id] != nil {
	// 	if userDb[user.Id].Email == user.Email {
	// 		errors.StatusBadRequestError(fmt.Sprintf("Email %v already exists", user.Email))
	// 	}
	// 	return errors.StatusBadRequestError(fmt.Sprintf("user %v already exists", user.Id))
	// }
	// userDb[user.Id] = user
	return nil
}

func (u *User) Update() *errors.RestErr {
	_, err := usersdb.Client.Exec(queryUpdateUser, u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	return nil
}

func (U *User) Delete() *errors.RestErr {
	_, err := usersdb.Client.Exec(queryDeleteUser, U.Id)
	if err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	return nil
}

func (u *User) FindUser(status string) ([]User, *errors.RestErr) {
	row, err := usersdb.Client.Query(queryFindUserByStatus, status)
	//usersdb.Client.Query()
	if err != nil {
		return nil, errors.StatusInternalServerError(err.Error())
	}
	results := make([]User, 0)
	for row.Next() {
		var user User
		if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.StatusInternalServerError(err.Error())
		}
		results = append(results, user)

		if len(results) == 0 {
			return nil, errors.StatusInternalServerError(err.Error())
		}
	}
	return results, nil
}
