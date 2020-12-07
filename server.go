package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tyshkorostyslav/test_golang_app/controllers/controller"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	v1 := router.Group("api/v1")
	{
		v1.POST("/user", Handlers.CreateUser)
		v1.POST("/resize/:id", Handlers.ResizePicture)
		v1.POST("/second_resize/:id", Handlers.SecondResizePicture)
		v1.GET("/all/:id", Handlers.GetRequestAllResizeObjs)
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// Start serving the application
	router.Run(":" + port)

}
