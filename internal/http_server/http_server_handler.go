package http_server

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	fmt.Println(r)
	_, err := fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
	if err != nil {
		return
	}
}