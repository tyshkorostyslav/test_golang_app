package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repositories "github.com/tyshkorostyslav/test_golang_app/repositories"
	services "github.com/tyshkorostyslav/test_golang_app/services"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		panic("couldn't get a database")
	}

	var user repositories.User
	c.Bind(&user)

	if user.Name != "" && user.HashedPword != "" {
		db.Create(&user)
		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

}

func GetRequestAllResizeObjs(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		panic("couldn't get a database")
	}
	id := c.Params.ByName("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	var resizingObjs []repositories.ResizingObj
	db.Where(map[string]interface{}{"user_id": user_id}).Find(&resizingObjs)
	// db.Find(&resizingObjs)

	c.JSON(200, resizingObjs)
}

func ResizePicture(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		panic("couldn't get a database")
	}
	id := c.Param("id")

	height := c.PostForm("height")
	width := c.PostForm("width")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	}
	log.Println(file.Filename)
	dst := "./upload/" + file.Filename
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	resized_url, err := services.CreateResizeObjInDB(
		db,
		dst,
		id,
		height,
		width,
		false,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"url":         dst,
			"resized_url": resized_url,
		})
	}

}

func SecondResizePicture(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get a database"})
		return
	}
	id := c.Param("id")
	img_id := c.PostForm("img_id")
	height := c.PostForm("height")
	width := c.PostForm("width")

	var resizingObjs []repositories.ResizingObj
	dst := "./upload/" + "resized_resized_" + img_id
	db.Where(map[string]interface{}{"user_id": id, "url": dst}).Find(&resizingObjs)
	if len(resizingObjs) == 0 {
		dst := "./upload/" + img_id
		resized_url, err := services.CreateResizeObjInDB(
			db,
			dst,
			id,
			height,
			width,
			true,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"url":         dst,
				"resized_url": resized_url,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Out of resizing limit"})
	}

}
