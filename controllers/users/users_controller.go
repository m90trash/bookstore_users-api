package users

import (
	"github.com/gin-gonic/gin"
	"github.com/m90trash/bookstore_users-api/domain/users"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())

	}

	c.JSON(http.StatusCreated, &user)

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//func SearchUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
