package api

import (
	"github.com/gin-gonic/gin"
)

type APIGroup struct {
	Engine *gin.Engine
	Group  *gin.RouterGroup
}

type IRoutable interface {
	ApplyRoutes() *APIGroup
}
