package main

import (
	"backend-jona-golang/config"
	"backend-jona-golang/databases"
	middleware "backend-jona-golang/midelware"
	migrate "backend-jona-golang/migration"
	router "backend-jona-golang/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.InitConfigEnv()

	// Koneksi ke database
	databases.Connect()
	migrate.Migrate()

	app := gin.Default()

	// midelware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.ErrorHandlingMiddleware())

	// Setup static file serving for images
	app.Static("/images", "./public/image-fitur")

	// Setup router
	router.RouterIndex(app)

	// Jalankan aplikasi di port 8080
	err := app.Run(config.APP_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
