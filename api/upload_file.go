package api

import (
	"HTTP_server/pkg"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
)

type ReceivedFile struct {
	Url string `json:"file"`
	Id  string
}

func UploadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	var p ReceivedFile
	err := json.NewDecoder(request.Body).Decode(&p)
	fmt.Println(p.Url)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("uploading File...")
	content, err := Download(p.Url)
	if err != nil {
		resp := make(map[string]string)
		resp["error"] = err.Error()
		jsonResp, _ := json.Marshal(resp)
		response.Write(jsonResp)
		return
	}
	fileName := pkg.GetFileNameFromUrl(p.Url)
	file, accessKey := pkg.Hashing(content)
	err = pkg.SaveToLocal(fileName, accessKey, 1024, file)
	fmt.Println(fileName, accessKey)
	fmt.Println("upload successfully!")
	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp := make(map[string]string)
		resp["error"] = err.Error()
		jsonResp, _ := json.Marshal(resp)
		response.Write(jsonResp)
		return
	}
	resp := make(map[string]string)
	resp["file_id"] = accessKey + ":" + fileName
	jsonResp, _ := json.Marshal(resp)
	response.Write(jsonResp)
	return
}
