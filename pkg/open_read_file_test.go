package pkg

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestOpenReadFileErrors(t *testing.T) {
	a, err := ReadFromLocal("hi", "tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
	fmt.Println(a, err)
	strings.EqualFold(err.Error(), "There was a problem in upload process")
}

func TestOpenReadFile(t *testing.T) {
	a, _ := ReadFromLocal("5b46015f09b4e16c857a3f8c5af8706b481036b5", "tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
	aa, _ := os.ReadFile("../data/5b46015f09b4e16c857a3f8c5af8706b481036b5/tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
	strings.EqualFold(string(a), string(aa))
}
