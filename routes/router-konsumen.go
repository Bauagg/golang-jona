package router

import (
	konsumencontrollers "backend-jona-golang/controlers/konsumen-controller"
	middleware "backend-jona-golang/midelware"

	"github.com/gin-gonic/gin"
)

func RouterKonsumen(app *gin.Engine) {
	router := app

	// Router Profile Konsumen
	router.GET("/api/jona/v-1/profile-konsumen", middleware.AuthMiddleware(), konsumencontrollers.ProfileKonsumen)

	// Router Pesanan Konsumen
	router.GET("/api/jona/v-1/pesanan-konsumen", middleware.AuthMiddleware(), konsumencontrollers.ListPesananKonsumen)
	router.GET("/api/jona/v-1/pesanan-konsumen/:id", middleware.AuthMiddleware(), konsumencontrollers.DetailPesananKonsumen)
	router.POST("/api/jona/v-1/pesanan-konsumen", middleware.AuthMiddleware(), konsumencontrollers.CreatePesanan)
	// Router Notofikasi Pembayaran
	router.POST("/api/jona/v-1/notifikasi-pembayaran", konsumencontrollers.NotifikasiPembayaran)
}
