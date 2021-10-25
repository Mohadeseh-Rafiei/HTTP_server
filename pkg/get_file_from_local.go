package pkg

import (
	"HTTP_server/internal"
	"fmt"
)

func GetFileFromLocal(file_id string) ([]byte, error) {
	directory, filename := SplitDirectoryAndFilename(file_id)
	fmt.Println("file:")
	fmt.Println(directory, filename)
	result, err := ReadFromLocal(directory, filename)
	if err != nil {
		return nil, internal.UnsuccessfulUpload
	}
	return result, nil
}
