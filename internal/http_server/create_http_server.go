package http_server

import (
	"HTTP_server/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Create() {
	_, err := fmt.Fprintf(os.Stdout, "Web Server started. Listening on localhost:8080\n")
	if err != nil {
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/", Handler)
	router.HandleFunc("/uploadFile", api.UploadFile).Methods("POST")
	router.HandleFunc("/downloadFile", api.DownloadFile).Methods("POST")
	fmt.Println("hi")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
