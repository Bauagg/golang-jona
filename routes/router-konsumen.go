package router

import (
	konsumencontrollers "backend-jona-golang/controlers/konsumen-controller"
	middleware "backend-jona-golang/midelware"

	"github.com/gin-gonic/gin"
)

func RouterKonsumen(app *gin.Engine) {
	router := app

	router.GET("/api/jona/v-1/pesanan-konsumen/:id", middleware.AuthMiddleware(), konsumencontrollers.DetailPesananKonsumen)
	router.POST("/api/jona/v-1/pesanan-konsumen", middleware.AuthMiddleware(), konsumencontrollers.CreatePesanan)
}
