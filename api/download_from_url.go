package api

import (
	"HTTP_server/internal"
	"io"
	"net/http"
)

func Download(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, internal.BadRequestError
	}
	return resp.Body, nil
}
