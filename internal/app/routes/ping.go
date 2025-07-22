package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getPingHandler handle the health check/ping endpoing
// @Summary 			Ping the Server
// @Description 	Simple health check endpoint
// @Tags 					Health
// @Produce 			plain
// @Success 			200 {string} string "pong"
// @Router 				/ping [get]
func getPingHandler(c *gin.Context) {
		c.String(http.StatusOK, "pong")
}

func GetPing(router *gin.RouterGroup) {
	router.GET("/ping", getPingHandler)
}
