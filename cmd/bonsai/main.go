// @title           Bonsai API
// @version         1.0
// @description     Backend service for Bonsai
// @host            localhost:8080
// @BasePath        /api
package main

import (
	"github.com/barely-sad-one/bonsai/internal/app/config"
	"github.com/barely-sad-one/bonsai/internal/app/server"
	"github.com/barely-sad-one/bonsai/internal/app/database"
)

func main() {
	config.InitConfig()
	database.InitDB()
	app := server.InitServer()
	app.Run()
}
