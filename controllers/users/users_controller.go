package users

import (
	"net/http"
	"strconv"

	"github.com/claudiocleberson/bookstore_oauth-shared/oauth"
	"github.com/claudiocleberson/bookstore_users-api/domain/users"
	"github.com/claudiocleberson/bookstore_users-api/services"
	"github.com/claudiocleberson/bookstore_utils-shared/utils/rest_err"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	//validate the request against the oauth shared library
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// //Certify the user inform the id
	// if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
	// 	err := rest_err.RestErr{
	// 		Code:    http.StatusUnauthorized,
	// 		Message: "resource no available",
	// 	}
	// 	c.JSON(err.Code, err)
	// 	return
	// }

	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(true))
		return
	}

	c.JSON(http.StatusOK, user.Marshal(oauth.IsPrivate(c.Request)))

}

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO - handle error
		restErr := rest_err.NewBadRequestError("invalid json body: " + err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {

		c.JSON(saveErr.Code, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshal(oauth.IsPrivate(c.Request)))
}

func UpdateUser(c *gin.Context) {

	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_err.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	var isPartial bool
	if c.Request.Method == http.MethodPatch {
		isPartial = true
	}

	user.Id = userId
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(oauth.IsPrivate(c.Request)))

}

func DeleteUser(c *gin.Context) {

	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func getUserId(userIdParam string) (int64, *rest_err.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		err := rest_err.NewBadRequestError("user id should be a number")

		return 0, err
	}
	return userId, nil
}

func SearchUser(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(oauth.IsPrivate(c.Request)))
}

func Login(c *gin.Context) {

	request := users.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_err.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.UsersService.UserLogin(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(oauth.IsPrivate(c.Request)))

}
