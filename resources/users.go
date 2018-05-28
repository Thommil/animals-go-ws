package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/api"
)

// Users Resource type
type Users api.Resource

// ApplyRoutes implements IRoutable interface
func (users *Users) ApplyRoutes() *Users {
	group := users.Engine.Group("/users")
	{
		group.GET("", users.findUser)
		group.GET("/:id", users.getUser)
		group.DELETE("/:id", users.deleteUser)
	}
	return users
}

func (users *Users) findUser(c *gin.Context) {
	c.String(http.StatusOK, "FIND users")
}

func (users *Users) getUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "GET %s", id)
}

func (users *Users) deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "DELETE %s", id)
}
