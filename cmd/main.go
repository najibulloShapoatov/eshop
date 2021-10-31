package main

import (
"eshop/internal/app"
_ "eshop/docs"
)

const (
	configFolder = "configs"
	configName = "main"
)

// @title Eshop API
// @version 1.0
// @description API Server for Eshop Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run(configFolder, configName)
}
