package controlers

import (
	"backend-jona-golang/databases"
	"backend-jona-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListDataCategory(ctx *gin.Context) {
	var data []models.CaegoryUtama

	err := databases.DB.
		Table("caegory_utamas").
		Where("fitur_id = ?", ctx.Param("id")).
		Preload("Fitur").
		Find(&data).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve CaegoryUtama",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data success",
		"data":    data,
	})
}

func ListDataCategoryAll(ctx *gin.Context) {
	var data []models.CaegoryUtama

	err := databases.DB.
		Table("caegory_utamas").
		Preload("Fitur").
		Find(&data).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve CaegoryUtama",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data success",
		"data":    data,
	})
}

func CreateCategoryUtama(ctx *gin.Context) {
	var input models.InputCaegoryUtama

	// Bind input from the request body
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Prepare the data to be inserted
	data := models.CaegoryUtama{
		FiturId:     input.FiturId,
		Nama:        input.Nama,
		Judul:       input.Judul,
		Description: input.Description,
	}

	// Create new category record in the database
	if err := databases.DB.Table("caegory_utamas").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create category",
		})
		return
	}

	// Return success response
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Category created successfully",
		"data":    data,
	})
}

func UpdateCategoryUtama(ctx *gin.Context) {
	var input models.InputCaegoryUtama
	var data models.CaegoryUtama

	// Bind input from the request body
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Find the existing category by ID
	if err := databases.DB.Table("caegory_utamas").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Category Utama with the specified ID was not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "An error occurred while retrieving the Category Utama: " + err.Error(),
			})

			return
		}
	}

	// Update the existing category fields
	data.Judul = input.Judul
	data.Nama = input.Nama
	data.Description = input.Description

	// Save the changes to the database
	if err := databases.DB.Table("caegory_utamas").Where("id = ?", ctx.Param("id")).Updates(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update Category Utama: " + err.Error(),
		})
		return
	}

	// Return success response for update
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Category Utama updated successfully",
		"data":    data,
	})
}

func DeleteCategoryUtama(ctx *gin.Context) {
	var data models.CaegoryUtama

	// Find the existing category by ID
	if err := databases.DB.Table("caegory_utamas").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Category Utama with the specified ID was not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "An error occurred while retrieving the Category Utama: " + err.Error(),
			})

			return
		}
	}

	// Delete the category
	if err := databases.DB.Table("caegory_utamas").Where("id = ?", ctx.Param("id")).Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to delete Category Utama: " + err.Error(),
		})
		return
	}

	// Return success response for deletion
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Category Utama deleted successfully",
	})
}
