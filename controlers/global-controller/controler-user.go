package controlers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	"backend-jona-golang/utils"
	"math/rand"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(ctx *gin.Context) {
	var input models.Users
	var otp models.OTP

	// Bind input JSON ke struct
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Validasi format email
	regexEmaill := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexEmaill.MatchString(input.Email) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid email format",
		})

		return
	}

	// Cek apakah email sudah terdaftar
	var existingUser models.Users
	err := databases.DB.Table("users").Where("email = ?", input.Email).First(&existingUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err == nil {
		// Jika tidak ada error dan user ditemukan, email sudah terdaftar
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Email sudah terdaftar.",
		})
		return
	}

	// Validasi panjang password
	if len(input.Password) <= 8 || input.Password == "" && len(input.KonfirmasiPassword) <= 8 || input.KonfirmasiPassword == "" {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "password kurang kuat.",
		})
		return
	}

	// validate password sama konfirm password harus sama
	if input.Password != input.KonfirmasiPassword {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Pastikan password dan konfirmasi password Anda sama.",
		})

		return
	}

	// Hash password
	input.Password = utils.HashPassword(input.Password)
	input.KonfirmasiPassword = utils.HashPassword(input.KonfirmasiPassword)

	// Simpan user baru ke database
	errCreate := databases.DB.Table("users").Create(&input).Error
	if errCreate != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to register user.",
		})
		return
	}

	// Generate a random 4-digit OTP
	rand.Seed(time.Now().UnixNano())
	randomOTP := rand.Intn(9000) + 1000

	// Save OTP to the database
	otp = models.OTP{
		NumberOtp: uint64(randomOTP),
		UserId:    uint64(input.ID),
		ExpiresAt: time.Now().Add(1 * time.Minute), // OTP expires in 1 minutes
	}

	errCreateOtp := databases.DB.Table("otps").Create(&otp).Error
	if errCreateOtp != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to generate OTP.",
		})
		return
	}

	// Send OTP email
	errSendEmail := utils.SendEmail(input.Email, uint64(randomOTP))
	if errSendEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to send OTP email.",
		})

		return
	}

	// Create Token
	token, err := utils.SignToken(uint64(input.ID), input.Email, string(input.Role))

	if err != nil {
		ctx.JSON(500, gin.H{ // Status 500 for Internal Server Error
			"error":   true,
			"message": "Failed to generate token.",
		})
		return
	}

	// Berhasil membuat user
	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "register success.",
		"datas":   input,
		"token":   token,
	})
}

func LoginUser(ctx *gin.Context) {
	var input models.InputLogin

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Cek apakah email sudah terdaftar
	var user models.Users
	err := databases.DB.Table("users").Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{ // Status 401 untuk Unauthorized
				"error":   true,
				"message": "Invalid email or password.",
			})
			return
		}

		ctx.JSON(500, gin.H{ // Status 500 untuk Internal Server Error
			"error":   true,
			"message": "Internal server error.",
		})
		return
	}

	// Verify password
	err = utils.VerifikasiHashPassword(input.Password, user.Password)
	if err != nil {
		ctx.JSON(401, gin.H{ // Status 401 for Unauthorized
			"error":   true,
			"message": "Invalid email or password.",
		})
		return
	}

	token, err := utils.SignToken(uint64(user.ID), user.Email, string(user.Role))

	if err != nil {
		ctx.JSON(500, gin.H{ // Status 500 for Internal Server Error
			"error":   true,
			"message": "Failed to generate token.",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomOTP := rand.Intn(9000) + 1000

	otp := models.OTP{
		NumberOtp: uint64(randomOTP),
		UserId:    uint64(user.ID),
		ExpiresAt: time.Now().Add(1 * time.Minute), // OTP expires in 5 minutes
	}

	errSendEmail := utils.SendEmail(input.Email, uint64(randomOTP))
	if errSendEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to send OTP email.",
		})

		return
	}

	if err := databases.DB.Table("otps").Where("user_id = ?", user.ID).Updates(otp).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update OTP.",
		})
		return
	}

	type Data struct {
		ID    uint64
		Email string
		Role  string
		Token string
	}

	data := Data{
		ID:    uint64(user.ID),
		Email: user.Email,
		Role:  string(user.Role),
		Token: token,
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "login success.",
		"data":    data,
	})
}

func CreateEmailOTP(ctx *gin.Context) {
	var input models.InputEmail
	var data models.Users

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	if err := databases.DB.Table("users").Where("email = ?", input.Email).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{ // Status 401 untuk Unauthorized
				"error":   true,
				"message": "Invalid email or password.",
			})
			return
		}

		ctx.JSON(500, gin.H{ // Status 500 untuk Internal Server Error
			"error":   true,
			"message": "Internal server error.",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomOTP := rand.Intn(9000) + 1000

	otp := models.OTP{
		NumberOtp: uint64(randomOTP),
		UserId:    uint64(data.ID),
		ExpiresAt: time.Now().Add(1 * time.Minute), // OTP expires in 5 minutes
	}

	errSendEmail := utils.SendEmail(input.Email, uint64(randomOTP))
	if errSendEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to send OTP email.",
		})

		return
	}

	if err := databases.DB.Table("otps").Where("user_id = ?", data.ID).Updates(otp).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update OTP.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OTP updated successfully.",
		"userId":  data.ID,
	})
}

func UpdatePassword(ctx *gin.Context) {
	var input models.InputPassword

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// validate password sama konfirm password harus sama
	if input.Password != input.KonfirmasiPassword {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Pastikan password dan konfirmasi password Anda sama.",
		})

		return
	}

	hashedPassword := utils.HashPassword(input.Password)

	if err := databases.DB.Table("users").Where("id = ?", ctx.Param("id")).Update("password", hashedPassword).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update the password: " + err.Error(),
		})
		return
	}

	// Return success response
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Password updated successfully.",
	})
}

func ValidatePassword(ctx *gin.Context) {
	var input models.InputValidatePassword
	var data models.Users

	user_id, _ := ctx.Get("userID")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
	}

	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "User not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve user profile",
			})

			return
		}
	}

	// Verify password
	err := utils.VerifikasiHashPassword(input.Password, data.Password)
	if err != nil {
		ctx.JSON(401, gin.H{ // Status 401 for Unauthorized
			"error":   true,
			"message": "Invalid password.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success konfirmasi password",
	})
}
