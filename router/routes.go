package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mediumAPI/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.POST("/text", handler.CreateText)
	}
}
