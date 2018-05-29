package users

import (
	"net/http"

	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
)

type users struct {
	group *gin.RouterGroup
}

// New create new users resource
func New(engine *gin.Engine) resource.Routable {
	users := &users{group: engine.Group("/users")}
	{
		users.group.GET("", users.findUser)
		users.group.GET("/:id", users.getUser)
		users.group.DELETE("/:id", users.deleteUser)
	}
	return users
}

// GetGroup implementation of Routable
func (users *users) GetGroup() *gin.RouterGroup {
	return users.group
}

func (users *users) findUser(c *gin.Context) {
	c.String(http.StatusOK, "FIND users")
}

func (users *users) getUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "GET %s", id)
}

func (users *users) deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "DELETE %s", id)
}
