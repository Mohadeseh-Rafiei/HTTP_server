package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Wanted struct {
	Url string `json:"file_id"`
}


func DownloadFile(response http.ResponseWriter, request *http.Request){
	request.ParseMultipartForm(10 * 1024 * 1024)
	var p Wanted
	err := json.NewDecoder(request.Body).Decode(&p)
	fmt.Println(p.Url)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("downloading...")

}
