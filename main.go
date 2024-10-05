package main

import (
	"backend-jona-golang/config"
	"backend-jona-golang/cronjob"
	"backend-jona-golang/databases"
	middleware "backend-jona-golang/midelware"
	migrate "backend-jona-golang/migration"
	router "backend-jona-golang/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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
	app.Static("/images-fitur", "./public/image-fitur")
	app.Static("/images-bank", "./public/image-bank")
	app.Static("/profile-user", "./public/profile-user")

	// Setup router
	router.RouterIndex(app)
	router.RouterKonsumen(app)

	// cronjob
	c := cron.New()

	// Menjalankan UpdateExpiredOrders setiap 1 menit
	if _, err := c.AddFunc("@every 1m", cronjob.UpdateExpiredOrders); err != nil {
		log.Fatalf("Error scheduling UpdateExpiredOrders: %v", err)
	}

	// Memulai cron
	c.Start()

	// Jalankan aplikasi di port 8080
	err := app.Run(config.APP_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
