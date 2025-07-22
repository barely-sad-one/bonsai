package server

import (
	"github.com/barely-sad-one/bonsai/internal/app/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}


func InitServer() *Server{
	var app *Server = &Server{}

	app.Router = gin.Default()

	routes.SetupRoutes(app.Router)
	return app
}

func (app *Server)Run() {
	app.Router.Run()
}
