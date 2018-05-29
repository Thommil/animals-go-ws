// Package animals defines resources used for /animals endpoint
package animals

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
)

type animals struct {
	group    *gin.RouterGroup
	database *mgo.Database
}

// New creates new Routable implementation for /animals resource
func New(engine *gin.Engine, database *mgo.Database) resource.Routable {
	animals := &animals{group: engine.Group("/animals"), database: database}
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
	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    http.StatusNotImplemented,
		"message": "Not implemented yet",
	})
}

func (animals *animals) getAnimals(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    http.StatusNotImplemented,
		"message": "Not implemented yet",
	})
}

func (animals *animals) deleteAnimals(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    http.StatusNotImplemented,
		"message": "Not implemented yet",
	})
}
