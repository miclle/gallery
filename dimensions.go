package imginfo

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Dimension return image width and height
func Dimension(imagePath string) (width, height int, err error) {
	var file *os.File

	file, err = os.Open(imagePath)
	if err != nil {
		return
	}

	defer file.Close()

	image, _, err := image.DecodeConfig(file)

	if err != nil {
		return
	}

	return image.Width, image.Height, nil
}
