package repositories

import "github.com/jinzhu/gorm"

type ResizingObj struct {
	gorm.Model
	OriginalURL  string           `db:"url" json:"url"`
	ResizedURL   string           `db:"resized_url" json:"resized_url"`
	ResizeParams ResizeParameters `gorm:"embedded"`
	UserID       int              `db:"user_id" json:"user_id"`
}

type User struct {
	gorm.Model
	Name         string `db:"name" json:"name"`
	HashedPword  string `db:"pword" json:"pword"`
	ResizingObjs []ResizingObj
}

type ResizeParameters struct {
	Height uint `db:"height" json:"height"`
	Width  uint `db:"width" json:"width"`
}
