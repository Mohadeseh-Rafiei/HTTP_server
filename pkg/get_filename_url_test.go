package pkg

import (
	"strings"
	"testing"
)

func TestGetFilename(t *testing.T) {
	fileUrl := "http://www.africau.edu/images/default/sample.pdf"
	filename := GetFileName(fileUrl)
	strings.EqualFold(filename, "sample.pdf")
}
