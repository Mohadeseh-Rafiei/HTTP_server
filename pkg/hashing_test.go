package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestHashing(t *testing.T) {
	resp, _ := http.Get("https://blog.photofeeler.com/wp-content/uploads/2017/09/tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
	_, hashed := Hashing(resp.Body)
	fmt.Println(hashed)
	strings.EqualFold(hashed, "5b46015f09b4e16c857a3f8c5af8706b481036b5")
}
