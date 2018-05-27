package animals

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Defines subroutes for /animals
func ApplyRoutes(router *gin.Engine) *gin.RouterGroup {
	animals := router.Group("/animals")
	{
		animals.POST("", create)
		animals.GET("", find)
	}
	return animals
}

func find(c *gin.Context) {
	c.String(http.StatusOK, "FIND animals")
}

func create(c *gin.Context) {
	c.String(http.StatusOK, "CREATE animals")
}
