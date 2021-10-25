package pkg

import (
	"path"
)

func GetFileName(url string) string {
	return path.Base(url)
}
