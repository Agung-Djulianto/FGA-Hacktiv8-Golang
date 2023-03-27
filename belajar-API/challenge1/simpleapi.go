package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pegawai struct {
	ID      string `json:"id"`
	Nama    string `json:"nama"`
	Umur    string `json:"umur"`
	Jabatan string `json:"jabatan"`
}

var DataPegawai = []Pegawai{}

func main() {
	router := gin.Default()

	router.GET("/all-pegawai", AllPegawai)
	router.GET("/get-pegawai/:pegawaiID", getPegawai)
	router.POST("/pegawai", CreatePegawai)
	router.PUT("/put-pegawai/:pegawaiID", UpdatePegawai)
	router.DELETE("delete-pegawai/:pegawaiID", DeletePegawai)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func AllPegawai(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DataPegawai)
}

func getPegawai(ctx *gin.Context) {
	pegawaiID := ctx.Param("pegawaiID")
	condition := false
	var dataPegawai Pegawai

	for i, pegawai := range DataPegawai {
		if pegawaiID == pegawai.ID {
			condition = true
			dataPegawai = DataPegawai[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("pegawai dengan id %v Tidak Ditemukan", pegawaiID),
		})
		return
	}
	ctx.JSON(http.StatusOK, dataPegawai)
}

func CreatePegawai(ctx *gin.Context) {
	var newPegawai Pegawai

	if err := ctx.ShouldBindJSON(&newPegawai); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newPegawai.ID = fmt.Sprintf("%d", len(DataPegawai)+1)
	DataPegawai = append(DataPegawai, newPegawai)

	ctx.JSON(http.StatusCreated, gin.H{
		"pegawai": newPegawai,
	})
}

func UpdatePegawai(ctx *gin.Context) {
	pegawaiID := ctx.Param("pegawaiID")
	condition := false
	var updatePegawai Pegawai

	if err := ctx.ShouldBindJSON(&updatePegawai); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, pegawai := range DataPegawai {
		if pegawaiID == pegawai.ID {
			condition = true
			DataPegawai[i] = updatePegawai
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_massage": "Data Tidak Ditemukan",
			"error_message": fmt.Sprintf("pegawai dengan id %v Tidak Ditemukan", pegawaiID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("karyawan dengan id %v telah diupdate", pegawaiID),
	})
}

func DeletePegawai(ctx *gin.Context) {
	pegawaiID := ctx.Param("pegawaiID")
	condition := false
	var pegawaiIndx int

	for i, pegawai := range DataPegawai {
		if pegawaiID == pegawai.ID {
			condition = true
			pegawaiIndx = i
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Tidak Ditemukan",
			"error_massage": fmt.Sprintf("Pegawai dengan id %v tidak Ditemukan", pegawaiID),
		})
		return
	}

	copy(DataPegawai[pegawaiIndx:], DataPegawai[pegawaiIndx+1:])
	DataPegawai[len(DataPegawai)-1] = Pegawai{}
	DataPegawai = DataPegawai[:len(DataPegawai)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("Pegawai dengan id %v berhasil terhapus", pegawaiID),
	})
}
