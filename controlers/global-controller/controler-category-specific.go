package controlers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"

	"github.com/gin-gonic/gin"
)

func ListCategorySpecific(ctx *gin.Context) {
	var data []models.CategorySpecific

	if err := databases.DB.Table("category_specifics").Where("id_sub_category = ?", ctx.Param("id")).Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Category Specific Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list category specific success",
		"data":    data,
	})
}

func CreateCategorySpecific(ctx *gin.Context) {
	var input models.CategorySpecific

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("category_specifics").Create(&input).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Category Specific Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create Category Specific success",
		"data":    input,
	})
}
