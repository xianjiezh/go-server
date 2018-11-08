package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "http server")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":12345", mux)
}

