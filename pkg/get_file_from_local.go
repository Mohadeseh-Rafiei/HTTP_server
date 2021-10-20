package pkg

import (
	"fmt"
)

func GetFileFromLocal(file_id string) []byte {
	directory, filename := SplitDirectoryAndFilename(file_id)
	fmt.Println("file:")
	fmt.Println(directory, filename)
	result := OpenReadFile(directory, filename)
	return result
}
