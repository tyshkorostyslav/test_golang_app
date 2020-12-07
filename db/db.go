package database

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	repositories "github.com/tyshkorostyslav/test_golang_app/repositories"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./db/data.db")

	// Error
	if err != nil {
		panic(err)
	}
	// Display SQL queries
	db.LogMode(true)
	// Creating the table
	if !db.HasTable(&repositories.User{}) {
		db.CreateTable(&repositories.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&repositories.User{})
	}

	if !db.HasTable(&repositories.ResizingObj{}) {
		db.CreateTable(&repositories.ResizingObj{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&repositories.ResizingObj{})
	}

	return db
}

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}
