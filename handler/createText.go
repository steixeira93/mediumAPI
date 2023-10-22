package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateText(c *gin.Context) {
	request := CreateTextRequest{}
	c.BindJSON(&request)

	if err := request.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
