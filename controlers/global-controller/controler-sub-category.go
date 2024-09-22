package controlers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListSubCategory(ctx *gin.Context) {
	var data []models.SubCategory

	err := databases.DB.Table("sub_categories").
		Where("id_category = ?", ctx.Param("id")).
		Preload("CaegoryUtama").
		Preload("CaegoryUtama.Fitur").
		Find(&data).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "No subcategories found for the given category ID",
			})
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve subcategories",
			})
		}
		return
	}

	// Jika data ditemukan, kembalikan response success dengan status 200
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Subcategories retrieved successfully",
		"data":    data,
	})
}

func ListSubCategoryAll(ctx *gin.Context) {
	var data []models.SubCategory

	if err := databases.DB.Table("sub_categories").Preload("CaegoryUtama").Preload("CaegoryUtama.Fitur").Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve subcategories",
		})
		return
	}

	// Jika data ditemukan, kembalikan response success dengan status 200
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Subcategories retrieved successfully",
		"data":    data,
	})
}

func CreateSubCategory(ctx *gin.Context) {
	var input models.InputSubCategory

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data := models.SubCategory{
		Nama:        input.Nama,
		Description: input.Description,
		IdCategory:  input.IdCategory,
		Harga:       input.Harga,
	}

	if err := databases.DB.Table("sub_categories").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create subcategory",
		})
		return
	}

	// Jika berhasil, kembalikan response success
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Subcategory created successfully",
		"data":    input,
	})
}

func UpdateSubCategory(ctx *gin.Context) {
	var input models.InputSubCategory
	var data models.SubCategory

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("sub_categories").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "No subcategories found for the given category ID",
			})
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve subcategories",
			})
		}
		return
	}

	data.Nama = input.Nama
	data.Harga = input.Harga
	data.Description = input.Description

	if err := databases.DB.Table("sub_categories").Where("id = ?", ctx.Param("id")).Updates(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update subcategory",
		})
		return
	}

	// Jika sukses, kembalikan response status 200 dengan pesan sukses
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Subcategory updated successfully",
		"data":    data,
	})
}

func DeleteSubCategory(ctx *gin.Context) {
	var data models.SubCategory

	if err := databases.DB.Table("sub_categories").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "No subcategories found for the given category ID",
			})
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve subcategories",
			})
		}
		return
	}

	if err := databases.DB.Table("sub_categories").Where("id = ?", ctx.Param("id")).Delete(&data).Error; err != nil {
		// Jika ada kesalahan saat penghapusan, kembalikan status 500
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to delete subcategory",
		})
		return
	}

	// Jika berhasil dihapus, kembalikan response status 200 dengan pesan sukses
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Subcategory deleted successfully",
	})
}
