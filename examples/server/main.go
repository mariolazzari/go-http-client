package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("GET /", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Ciao Mario!")
}
