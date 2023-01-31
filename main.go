package main

import (
	"api/app/config"
	"api/app/lib"
	"api/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
}

// @title Project Title
// @version 1.0.0
// @description Project Name Description
// @termsOfService https://gospecs.monstercode.net/en/guide/tnc.html
// @contact.name Developer
// @contact.email developer.team@gmail.com
// @host localhost:9000
// @schemes http
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
