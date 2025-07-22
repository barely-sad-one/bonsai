package main

import (
	"github.com/barely-sad-one/bonsai/internal/app/config"
	"github.com/barely-sad-one/bonsai/internal/app/server"
)

func main() {
	config.InitConfig()
	app := server.InitServer()
	app.Run()
}
