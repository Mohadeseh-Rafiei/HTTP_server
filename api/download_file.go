package api

import (
	"HTTP_server/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

type Wanted struct {
	Id string `json:"file_id"`
}

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	fmt.Println("hi")
	var p Wanted
	json.NewDecoder(request.Body).Decode(&p)
	fmt.Println("downloading...")
	file_bytes := pkg.GetFileFromLocal(p.Id)
	response.Write(file_bytes)
}
