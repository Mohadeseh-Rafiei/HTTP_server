package api

import (
	"HTTP_server/internal"
	"HTTP_server/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

type DownloadedFile struct {
	Id string `json:"file_id"`
}

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	var p DownloadedFile
	json.NewDecoder(request.Body).Decode(&p)
	fmt.Println("downloading...")
	fileBytes, err := pkg.GetFileFromLocal(p.Id)
	_, err = response.Write(fileBytes)
	if err != nil {
		resp := make(map[string]string)
		resp["error"] = internal.UnsuccessfulUpload.Error()
		jsonResp, _ := json.Marshal(resp)
		response.Write(jsonResp)
		return
	}
}
