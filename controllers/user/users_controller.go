package user

import (
	"net/http"
	"strconv"

	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/services"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//these functions are entry point of our functions

//in this function we are handling that tries or attempt to create user
func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.StatusBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	Result, saveErr := services.Userservice.CreateUser(user)
	if saveErr != nil {

		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, Result.Marshal(c.GetHeader("X-Public") == "true"))
}

//This function is used when getting user from database
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.StatusBadRequestError("invalid User ID")
		c.JSON(err.Status, err)
		return
	}
	Result, getErr := services.Userservice.GetUser(userId)
	if getErr != nil {

		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, Result.Marshal(c.GetHeader("X-Public") == "true"))

}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.StatusBadRequestError("invalid User ID")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.StatusBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	Result, err := services.Userservice.UpdateUSer(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, Result.Marshal(c.GetHeader("X-Public") == "true"))

}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.StatusBadRequestError("invalid User ID")
		c.JSON(err.Status, err)
		return
	}

	err := services.Userservice.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "success"})

}

func FindUserStatus(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Userservice.FindByStatus(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result := make([]interface{}, len(*users))
	for i, v := range *users {
		result[i] = v.Marshal(c.GetHeader("X-Public") == "true")
	}
	c.JSON(http.StatusOK, result)
}
