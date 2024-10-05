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
	router.PUT("/api/jona/v-1/profile-konsumen", middleware.AuthMiddleware(), konsumencontrollers.UpdateProfileKonsumen)

	// Router Pesanan Konsumen
	router.GET("/api/jona/v-1/pesanan-konsumen", middleware.AuthMiddleware(), konsumencontrollers.ListPesananKonsumen)
	router.GET("/api/jona/v-1/pesanan-konsumen/:id", middleware.AuthMiddleware(), konsumencontrollers.DetailPesananKonsumen)
	router.POST("/api/jona/v-1/pesanan-konsumen", middleware.AuthMiddleware(), konsumencontrollers.CreatePesananBersihBersih)

	// Router Notofikasi Pembayaran
	router.POST("/api/jona/v-1/notifikasi-pembayaran", konsumencontrollers.NotifikasiPembayaran)

	// Router Address Tujuan
	router.GET("/api/jona/v-1/address-tujuan/:id", middleware.AuthMiddleware(), konsumencontrollers.DetailAddressTujuan)
	router.POST("/api/jona/v-1/address-tujuan", middleware.AuthMiddleware(), konsumencontrollers.CreateAddressTujuan)
	router.PUT("/api/jona/v-1/address-tujuan/:id", middleware.AuthMiddleware(), konsumencontrollers.UpdateAddressTujuan)
}
