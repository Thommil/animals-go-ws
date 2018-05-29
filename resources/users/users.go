// Package users defines resources used for /users endpoint
package users

import (
	"net/http"

	"github.com/thommil/animals-go-common/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
)

type users struct {
	group    *gin.RouterGroup
	database *mgo.Database
}

// New creates new Routable implementation for /users resource
func New(engine *gin.Engine, database *mgo.Database) resource.Routable {
	users := &users{group: engine.Group("/users"), database: database}
	{
		users.group.GET("", users.findUser)
		users.group.GET("/:id", users.getUser)
		users.group.DELETE("/:id", users.deleteUser)
	}
	return users
}

// GetGroup implementation of resource.Routable
func (users *users) GetGroup() *gin.RouterGroup {
	return users.group
}

func (users *users) findUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    http.StatusNotImplemented,
		"message": "Not implemented yet",
	})
}

func (users *users) getUser(c *gin.Context) {
	user := &model.User{}
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		if err := users.database.C("user").FindId(bson.ObjectIdHex(id)).One(user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid user ID",
		})
	}
}

func (users *users) deleteUser(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		if err := users.database.C("user").RemoveId(bson.ObjectIdHex(id)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
		} else {
			c.Status(http.StatusOK)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid user ID",
		})
	}
}
