package users

import (
	"net/http"
	"strconv"

	"github.com/dbhushan9/bookstore_users-api/domain/users"
	"github.com/dbhushan9/bookstore_users-api/services"
	"github.com/dbhushan9/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	var user users.User
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
