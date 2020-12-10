package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	controllers "github.com/tyshkorostyslav/test_golang_app/controllers/v1"
	database "github.com/tyshkorostyslav/test_golang_app/db"
)

var router *gin.Engine

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_user := os.Getenv("DB_USER")
	pword := os.Getenv("DB_PWORD")
	db_addr := os.Getenv("DB_ADDR")
	db_name := os.Getenv("DB_NAME")
	db := database.InitDb(db_user, pword, db_addr, db_name)
	defer db.Close()

	router = gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Use(database.ApiMiddleware(db))

	v1 := router.Group("api/v1")
	{
		v1.POST("/user", controllers.CreateUser)
		v1.POST("/resize/:id", controllers.ResizePicture)
		v1.POST("/second_resize/:id", controllers.SecondResizePicture)
		v1.GET("/all/:id", controllers.GetRequestAllResizeObjs)
	}

	port := os.Getenv("PORT")

	// Start serving the application
	router.Run(":" + port)

}
