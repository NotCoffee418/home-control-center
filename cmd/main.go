package main

import (
	"github.com/NotCoffee418/home-control-center/internal/config"
	"github.com/NotCoffee418/home-control-center/internal/db"
	"github.com/NotCoffee418/home-control-center/internal/web"
)

func main() {
	// Preload config
	_ = config.GetConfig()

	// Prepare database
	db.InitializeDatabase()

	// Start web server
	web.StartWebServer()
}
