package BenchMarks

import (
	"HTTP_server/api"
	"HTTP_server/pkg"
	"fmt"
	"os"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	content, _ := api.Download("https://blog.photofeeler.com/wp-content/uploads/2017/09/tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg")
	file, accessKey := pkg.Hashing(content)
	for i := 0; i < b.N; i++ {
		filePath := "./data/BenchMarkTests/UploadTest" + "/" + accessKey + "/" + string(rune(i))
		outFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		pkg.DoSaving(2, 1024*i, file, outFile)
	}
}
