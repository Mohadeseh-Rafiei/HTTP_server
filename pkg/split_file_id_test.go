package pkg

import (
	"strings"
	"testing"
)

func TestSplitDirectoryAndFilename(t *testing.T) {
	file_id := "5b46015f09b4e16c857a3f8c5af8706b481036b5:tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg"
	a, b := SplitDirectoryAndFilename(file_id)
	strings.EqualFold(a, "5b46015f09b4e16c857a3f8c5af8706b481036b5")
	strings.EqualFold(b, "tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
}
