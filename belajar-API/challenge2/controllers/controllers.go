package controllers

import (
	"challenge2/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InputBook struct {
	Titel     string `json:"title"`
	Author    string `json:"author"`
	Deskripsi string `json:"deskripsi"`
}

// fungsi untuk mendapatkan semua data buku
func GetAllBuku(ctx *gin.Context) {
	var results = []models.BOOK{}

	sqlStatement := "SELECT id, title, author, deskripsi FROM book"

	rows, err := ctx.MustGet("db").(*sql.DB).Query(sqlStatement)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.BOOK{}

		err = rows.Scan(&book.ID, &book.Titel, &book.Author, &book.Deskripsi)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		results = append(results, book)
	}

	if err = rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}

// fungsi untuk mendapatkan data buku berdasarkan ID primary key
func GetBukuId(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var dataBuku = models.BOOK{}
	var sqlQuery = "SELECT * FROM book WHERE id = $1"

	err := db.QueryRow(sqlQuery, bookID).Scan(&dataBuku.ID, &dataBuku.Titel, &dataBuku.Author, &dataBuku.Deskripsi)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Tidak Ditemukan",
			"error_massage": fmt.Sprintf("Buku dengan id %v tidak ditemukan", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, dataBuku)
}

// fungsi untuk membuat data buku baru
func CreateBuku(ctx *gin.Context) {

	var newBook InputBook
	var book = models.BOOK{}
	var db = ctx.MustGet("db").(*sql.DB)

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	sqlQuery := `INSERT INTO book (title,author,deskripsi)
					VALUES($1,$2,$3)
					RETURNING *
				`

	err := db.QueryRow(sqlQuery, newBook.Titel, newBook.Author, newBook.Deskripsi).Scan(&book.ID, &book.Titel, &book.Author, &book.Deskripsi)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return

	}
	ctx.JSON(http.StatusCreated, book)
}

// fungsi untuk update data buku berdasrkan ID
func PutBuku(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var buku models.BOOK
	var sqlStatement = `
		UPDATE book
		SET title = $2, author = $3, deskripsi = $4
		WHERE id = $1;
	`
	var whereid int
	var sqlStatementSelectId = `
		SELECT id FROM book WHERE id = $1
	`
	err := db.QueryRow(sqlStatementSelectId, bookID).
		Scan(&whereid)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Tidak Ditemukan",
			"error_message": fmt.Sprintf("Buku dengan id %v tidak ditemukan", bookID),
		})
		return
	}

	var input InputBook
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	buku.ID = bookID
	buku.Titel = input.Titel
	buku.Author = input.Author
	buku.Deskripsi = input.Deskripsi

	res, err := db.Exec(sqlStatement, buku.ID, buku.Titel, buku.Author, buku.Deskripsi)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, err := res.RowsAffected()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	if count > 0 {
		ctx.JSON(http.StatusOK, buku)
	}
}

// fungsi untuk menghapus data buku bersarkan ID
func DelBuku(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var whereid int
	var sqlQuery = ` SELECT id FROM book WHERE id = $1`

	err := db.QueryRow(sqlQuery, bookID).Scan(&whereid)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Tidak Ditemukan",
			"error_massage": fmt.Sprintf("buku dengan id %v tidak ditemukan", bookID),
		})
		return
	}

	var deleteSql = `DELETE FROM book WHERE id = $1;`

	res, err := db.Exec(deleteSql, bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err})
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("Buku dengan id %v terhapus", bookID),
	})
}
