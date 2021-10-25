package pkg

import (
	"path"
)

func GetFileNameFromUrl(url string) string {
	return path.Base(url)
}
