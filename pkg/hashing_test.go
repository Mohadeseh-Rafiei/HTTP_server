package pkg

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHashing(t *testing.T) {
	resp, _ := http.Get("http://www.africau.edu/images/default/sample.pdf")
	contetnt, hashed := Hashing(resp.Body)
	fmt.Println(hashed)
	fmt.Println(contetnt)
}
