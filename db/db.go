package database

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	repositories "github.com/tyshkorostyslav/test_golang_app/repositories"
)

func InitDb(db_user string, pword string, db_addr string, db_name string) *gorm.DB {
	// Openning file
	dsn := db_user + ":" + pword + "@" + db_addr + "/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

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
