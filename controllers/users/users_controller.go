package users

import (
	"fmt"
	"net/http"

	"github.com/marksfu1/bookstore_users-api/controllers/users/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/marksfu1/bookstore_users-api/domain/users"
	"github.com/marksfu1/bookstore_users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Printf("In Createuser. user: %v\n", user)

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

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}
