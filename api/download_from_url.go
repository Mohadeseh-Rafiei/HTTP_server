package api

import (
	"io"
	"net/http"
)

func Download(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	return resp.Body
}
