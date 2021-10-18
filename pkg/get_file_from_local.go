package pkg

import "fmt"

func GetFileFromLocal(file_id string) {
	directory, filename := SplitDirectoryAndFilename(file_id)
	fmt.Println(directory, filename)
}
