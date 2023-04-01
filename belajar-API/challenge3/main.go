package main

import (
	"challenge3/database"
	"challenge3/models"
	"challenge3/routers"
)

func main() {
	db := database.ReadDB()
	db.AutoMigrate(&models.Buku{})

	y := routers.ReadRouters(db)
	y.Run(":8080")

}
