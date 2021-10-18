package pkg

import (
	"path"
)

func GetFilename(url string) string  {
	return path.Base(url)
}