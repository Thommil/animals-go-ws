// Package users defines resources used for /users endpoint
package users

import (
	"net/http"

	"github.com/thommil/animals-go-common/model"

	"github.com/globalsign/mgo"
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

func (users *users) getMe(c *gin.Context) {
	if user, ok := c.Get("user"); ok {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Unable to recover session",
		})
	}
}

func (users *users) getUser(c *gin.Context) {
	if user, err := model.FindUserByID(users.database, c.Param("id")); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
	}
}

func (users *users) deleteUser(c *gin.Context) {
	if err := model.DeleteUserByID(users.database, c.Param("id")); err == nil {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
	}
}
