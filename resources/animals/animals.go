// Package animals defines resources used for /animals endpoint
package animals

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
)

type animals struct {
	group      *gin.RouterGroup
	collection *mgo.Collection
}

// New creates new Routable implementation for /animals resource
func New(engine *gin.Engine, mongo *mgo.Session) resource.Routable {
	animals := &animals{group: engine.Group("/animals"), collection: mongo.DB("").C("animal")}
	{
		animals.group.GET("", animals.findAnimals)
		animals.group.GET("/:id", animals.getAnimals)
		animals.group.DELETE("/:id", animals.deleteAnimals)
	}
	return animals
}

// GetGroup implementation of resource.Routable
func (animals *animals) GetGroup() *gin.RouterGroup {
	return animals.group
}

func (animals *animals) findAnimals(c *gin.Context) {
	c.String(http.StatusOK, "FIND animals")
}

func (animals *animals) getAnimals(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "GET %s", id)
}

func (animals *animals) deleteAnimals(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "DELETE %s", id)
}
