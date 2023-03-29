package routers

import (
	"challenge2/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.GET("/buku", controllers.GetAllBuku)
	router.GET("/buku/:bookID", controllers.GetBukuId)
	router.POST("/buku", controllers.CreateBuku)
	router.PUT("/buku/:bookID", controllers.PutBuku)
	router.DELETE("/buku/:bookID", controllers.DelBuku)

	return router
}
