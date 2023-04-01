package controllers

import (
	"challenge3/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BukuInput struct {
	JudulBuku string `json:"judul_buku"`
	Author    string `json:"author"`
	Kategori  string `json:"kategori"`
}

type InputUpdate struct {
	JudulBuku string    `json:"judul_buku"`
	Author    string    `json:"author"`
	Kategori  string    `json:"kategori"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllBuku(ctx *gin.Context) {
	var DataBuku []models.Buku
	db := ctx.MustGet("db").(*gorm.DB)
	db.Find(&DataBuku)
	ctx.JSON(http.StatusOK, DataBuku)
}

func GetBukuId(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	IdBuku := ctx.Param("IdBuku")
	var DataBuku models.Buku

	err := db.Where("id = ?", IdBuku).First(&DataBuku).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data tidak ditemukan",
			"error_massage": fmt.Sprintf("buku dengan id %v tidak ditemukan", IdBuku),
		})
		return
	}
	ctx.JSON(http.StatusOK, DataBuku)

}

func CreateBuku(ctx *gin.Context) {
	var bukuBaru BukuInput
	db := ctx.MustGet("db").(*gorm.DB)

	if err := ctx.ShouldBindJSON(&bukuBaru); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	bukuu := models.Buku{
		JudulBuku: bukuBaru.JudulBuku,
		Author:    bukuBaru.Author,
		Kategori:  bukuBaru.Kategori,
	}

	err := db.Create(&bukuu).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, bukuu)
}

func PutBuku(ctx *gin.Context) {

	IdBuku := ctx.Param("IdBuku")
	db := ctx.MustGet("db").(*gorm.DB)

	var bukuu models.Buku
	var input InputUpdate

	var UpdateBuku models.Buku
	UpdateBuku.JudulBuku = input.JudulBuku
	UpdateBuku.Author = input.Author
	UpdateBuku.Kategori = input.Kategori
	UpdateBuku.UpdatedAt = time.Now()

	err := db.Model(&bukuu).Where("id = ?", IdBuku).Updates(&UpdateBuku).Error

	if err := db.Where("id = ?", IdBuku).First(&bukuu).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data tidak ditemukan",
			"error_massage": fmt.Sprintf("buku dengan id %v tidak ditemukan", IdBuku),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, bukuu)
}

func DelBuku(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	IdBuku := ctx.Param("IdBuku")
	buku := models.Buku{}

	err := db.Where("id = ?", IdBuku).Delete(&buku).Error

	var bukuId models.Buku

	if err := db.Where("id = ?", IdBuku).First(&bukuId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  " data tidak ditemukan",
			"error_massage": fmt.Sprintf("buku dengan id %v tidak ditemukan", bukuId),
		})
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "buku berhasil dihapus",
	})
}
