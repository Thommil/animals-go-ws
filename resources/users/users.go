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
	group      *gin.RouterGroup
	collection *mgo.Collection
}

// New creates new Routable implementation for /users resource
func New(engine *gin.Engine, mongo *mgo.Session) resource.Routable {
	users := &users{group: engine.Group("/users"), collection: mongo.DB("").C("user")}
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
	c.String(http.StatusOK, "FIND users")
}

func (users *users) getUser(c *gin.Context) {
	user := &model.User{}
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		if err := users.collection.FindId(bson.ObjectIdHex(id)).One(user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err,
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
		if err := users.collection.RemoveId(bson.ObjectIdHex(id)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err,
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
