package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/api"
)

// Animals Resource type
type Animals api.Resource

// ApplyRoutes implements IRoutable interface
func (animals *Animals) ApplyRoutes() *Animals {
	animals.Group = animals.Engine.Group("/animals")
	{
		animals.Group.GET("", animals.findUser)
		animals.Group.GET("/:id", animals.getUser)
		animals.Group.DELETE("/:id", animals.deleteUser)
	}
	return animals
}

func (animals *Animals) findUser(c *gin.Context) {
	c.String(http.StatusOK, "FIND Animals")
}

func (animals *Animals) getUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "GET %s", id)
}

func (animals *Animals) deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "DELETE %s", id)
}