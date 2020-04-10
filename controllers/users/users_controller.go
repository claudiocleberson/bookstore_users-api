package users

import (
	"net/http"
	"strconv"

	"github.com/claudiocleberson/bookstore_users-api/domain/users"
	"github.com/claudiocleberson/bookstore_users-api/services"
	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := rest_err.NewBadRequestError("user id should be a number")
		c.JSON(err.Code, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO - handle error
		restErr := rest_err.NewBadRequestError("invalid json body: " + err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {

		c.JSON(saveErr.Code, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
