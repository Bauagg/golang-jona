package router

import (
	"backend-jona-golang/controlers"
	middleware "backend-jona-golang/midelware"

	"github.com/gin-gonic/gin"
)

func RouterIndex(app *gin.Engine) {
	router := app

	// Router untuk table User
	router.POST("/api/jona/v-1/register", controlers.RegisterUser)
	router.POST("/api/jona/v-1/login", controlers.LoginUser)

	// Router untuk OTP
	router.GET("/api/jona/v-1/otp", middleware.AuthMiddleware(), controlers.SendEmailOtp)
	router.POST("/api/jona/v-1/otp", middleware.AuthMiddleware(), controlers.VerifyOTP)

	// Router untuk Address
	router.GET("/api/jona/v-1/address", middleware.AuthMiddleware(), controlers.GetAddress)
	router.GET("/api/jona/v-1/address/:id", middleware.AuthMiddleware(), controlers.DetailAddress)
	router.POST("/api/jona/v-1/address", middleware.AuthMiddleware(), controlers.CreateAddress)
	router.PUT("/api/jona/v-1/address/:id", middleware.AuthMiddleware(), controlers.UpdateAddress)
	router.DELETE("/api/jona/v-1/address/:id", middleware.AuthMiddleware(), controlers.DeleteAddress)

	// Router untuk Fitur Jona
	router.GET("/api/jona/v-1/fitur-jona", controlers.GetFiturJona)
	router.POST("/api/jona/v-1/fitur-jona", controlers.CreateFitur)
	router.PUT("/api/jona/v-1/fitur-jona/:id", controlers.UpdateFiturJona)
	router.DELETE("/api/jona/v-1/fitur-jona/:id", controlers.DeleteFiturJona)
}
