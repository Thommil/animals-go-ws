package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Mongo : https://godoc.org/github.com/globalsign/mgo

// Defines subroutes for /users
func ApplyRoutes(router *gin.Engine) *gin.RouterGroup {
	users := router.Group("/users")
	{
		users.POST("", createUser)
		users.GET("", findUser)
		users.GET("/:id", getUser)
		users.DELETE("/:id", deleteUser)
	}
	return users
}

func findUser(c *gin.Context) {
	c.String(http.StatusOK, "FIND users")
}

func createUser(c *gin.Context) {
	c.String(http.StatusOK, "CREATE users")
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "GET %s", id)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "DELETE %s", id)
}
