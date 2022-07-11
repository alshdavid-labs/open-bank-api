package pinpad

import (
	"errors"
	"image"
)

func FindDigit(img *image.RGBA) (int, error) {
	for pos, knownImg := range numbers {
		diff, err := compare(knownImg, img)
		if err != nil {
			return 0, errors.New("UnableToCompare")
		}
		if diff < 1500 {
			return pos, nil
		}
	}
	return 0, errors.New("DigitNotFound")
}

func Base64ToRGBA(s string) (*image.RGBA, error) {
	v, err := b64ToImage(s)
	if err != nil {
		return nil, err
	}
	return imageToRGBA(v), nil
}
