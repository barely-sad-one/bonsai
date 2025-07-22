package server

import (
	"github.com/barely-sad-one/bonsai/internal/app/config"
	"github.com/barely-sad-one/bonsai/internal/app/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}


func InitServer() *Server{
	var app *Server = &Server{}

	gin.SetMode(config.AppConfig.Mode)
	app.Router = gin.New()

	routes.SetupRoutes(app.Router)
	return app
}

func (app *Server)Run() {
	app.Router.Run()
}
