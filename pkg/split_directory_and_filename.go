package pkg

import "strings"

func SplitDirectoryAndFilename(fileId string) (string, string) {
	splitedString := strings.Split(fileId, ":")
	return splitedString[0], splitedString[1]
}
