package users

import (
	"github.com/gin-gonic/gin"
	"github.com/m90trash/bookstore_users-api/domain/users"
	"github.com/m90trash/bookstore_users-api/services"
	"github.com/m90trash/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {

		restError := errors.NewBadRequestError("invalid json")

		c.JSON(restError.Status, restError)
		return

	}

	saveResult, saveError := services.CreateUser(user)

	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, saveResult)

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//func SearchUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
