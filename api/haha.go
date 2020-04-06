package handler

import (
	"fmt"
	"net/http"
)

func HaHa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>test.go</h1>")
}
