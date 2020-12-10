package services

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/tyshkorostyslav/test_golang_app/repositories"
	"github.com/tyshkorostyslav/test_golang_app/utils"
)

func CreateResizeObjInDB(db *gorm.DB, dst string, id string, height_string string, width_string string, second_resize bool) (resized_url string, err error) {

	height, err := strconv.Atoi(height_string)
	if err != nil {
		return "", err
	}
	width, err := strconv.Atoi(width_string)
	if err != nil {
		return "", err
	}
	dst_split := strings.Split(dst, "/")

	resized_url = "resized_" + dst_split[len(dst_split)-1]
	if second_resize {
		resized_url = "resized_" + resized_url
	}
	resized_url = "./upload/" + resized_url
	utils.Resize(dst, uint(height), uint(width), resized_url)

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	input := repositories.ResizingObj{
		OriginalURL: dst,
		ResizedURL:  resized_url,
		UserID:      user_id,
		ResizeParams: repositories.ResizeParameters{
			Height: uint(height),
			Width:  uint(width),
		},
	}

	db.Create(&input)

	return resized_url, nil
}
