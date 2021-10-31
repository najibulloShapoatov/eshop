package main

import (
	_ "eshop/docs"
	"eshop/internal/app"
)

const (
	configFolder = "configs"
	configName   = "main"
)

// @title Eshop API
// @version 1.0
// @description API Server for Eshop Application

// @host localhost:8383
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run(configFolder, configName)
}
