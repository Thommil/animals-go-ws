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
	group := animals.Engine.Group("/animals")
	{
		group.GET("", animals.findUser)
		group.GET("/:id", animals.getUser)
		group.DELETE("/:id", animals.deleteUser)
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
