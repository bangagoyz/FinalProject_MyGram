package main

import (
	"finalProject/database"
	"finalProject/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	routers := gin.Default()

	database.StartDB()
	db := database.GetDB()
	router.StartApp(routers, db)

	routers.Run(":" + os.Getenv("PORT"))
}
