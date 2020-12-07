package utils

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func Resize(image string, height uint, width uint, resized_file_name string) (err error) {
	file, err := os.Open(image)
	if err != nil {
		return err
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)
	out, err := os.Create(resized_file_name)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return nil
}
