package BenchMarks

import (
	"HTTP_server/api"
	"HTTP_server/pkg"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	url := "https://blog.photofeeler.com/wp-content/uploads/2017/09/tinder-photo-size-tinder-picture-size-tinder-aspect-ratio-image-dimensions-crop.jpg"
	content, _ := api.Download(url)
	fileName := path.Base(url)
	fileInfo := strings.Split(fileName, ".")
	file, _ := pkg.Hashing(content)
	for i := 1; i < 5; i++ {
		filePath := "../data/BenchMarkTests/UploadTest" + "/" + fileInfo[0] + strconv.Itoa(i) + "." + fileInfo[1]
		outFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		pkg.SaveByChunk(2, 1024*i, file, outFile)
	}
}
