package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPingHandler(c *gin.Context) {
		c.String(http.StatusOK, "pong")
}

func GetPing(router *gin.RouterGroup) {
	router.GET("/ping", getPingHandler)
}
