package router

import (
	controlers "backend-jona-golang/controlers/global-controller"
	middleware "backend-jona-golang/midelware"

	"github.com/gin-gonic/gin"
)

func RouterIndex(app *gin.Engine) {
	router := app

	// Router untuk table User
	router.POST("/api/jona/v-1/register", controlers.RegisterUser)
	router.POST("/api/jona/v-1/login", controlers.LoginUser)

	// Router lupa Password
	router.PUT("/api/jona/v-1/email", controlers.CreateEmailOTP)
	router.PUT("/api/jona/v-1/password", controlers.UpdatePassword)

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

	// Router untuk Category Utama
	router.GET("/api/jona/v-1/category-utama", controlers.ListDataCategoryAll)
	router.GET("/api/jona/v-1/category-utama/:id", controlers.ListDataCategory)
	router.POST("/api/jona/v-1/category-utama", controlers.CreateCategoryUtama)
	router.PUT("/api/jona/v-1/category-utama/:id", controlers.UpdateCategoryUtama)
	router.DELETE("/api/jona/v-1/category-utama/:id", controlers.DeleteCategoryUtama)

	//  Router untuk Sub category
	router.GET("/api/jona/v-1/sub-category", controlers.ListSubCategoryAll)
	router.GET("/api/jona/v-1/sub-category/:id", controlers.ListSubCategory)
	router.POST("/api/jona/v-1/sub-category", controlers.CreateSubCategory)
	router.PUT("/api/jona/v-1/sub-category/:id", controlers.UpdateSubCategory)
	router.DELETE("/api/jona/v-1/sub-category/:id", controlers.DeleteSubCategory)

	// Router untuk Daftar Bank
	router.GET("/api/jona/v-1/daftar-bank", controlers.ListDaftarBankAll)
	router.POST("/api/jona/v-1/daftar-bank", controlers.CreateDaftarBank)
	router.PUT("/api/jona/v-1/daftar-bank/:id", controlers.UpdateDaftarBank)
	router.DELETE("/api/jona/v-1/daftar-bank/:id", controlers.DeleteDaftarBank)
}
