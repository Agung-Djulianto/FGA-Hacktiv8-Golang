package routers

import (
	"challenge3/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadRouters(db *gorm.DB) *gin.Engine {
	x := gin.Default()
	x.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	x.GET("/buku", controllers.GetAllBuku)
	x.GET("/buku/:IdBuku", controllers.GetBukuId)
	x.POST("/buku", controllers.CreateBuku)
	x.PUT("/buku/:IdBuku", controllers.PutBuku)
	x.DELETE("/buku/:IdBuku", controllers.DelBuku)

	return x
}
