package internal

import "errors"

var (
	BadRequestError      = errors.New("bad request")
	UnsuccessfulDownload = errors.New("There was a problem in download process")
	UnsuccessfulUpload   = errors.New("There was a problem in upload process")
)
