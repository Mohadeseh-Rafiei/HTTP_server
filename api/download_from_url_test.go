package api

import "testing"

func TestDownload(t *testing.T) {
	fileUrl := "http://www.africau.edu/images/default/sample.pdf"
	Download(fileUrl)
}