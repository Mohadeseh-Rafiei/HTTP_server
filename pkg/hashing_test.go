package pkg

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHashing(t *testing.T) {
	resp, _ := http.Get("http://www.africau.edu/images/default/sample.pdf")
	hashed := Hashing(resp.Body)
	fmt.Println(hashed)
}
