package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler(...)")
	defer log.Println("............")
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
