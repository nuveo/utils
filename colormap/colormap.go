package colormap

import (
	"errors"
	"image/color"

	"golang.org/x/image/colornames"
)

func GetNameByRGB(r, g, b float64) (string, error) {
	rgb := color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}
	for key, value := range colornames.Map {
		if value == rgb {
			return key, nil
		}
	}

	return "", errors.New("Color not found.")
}
