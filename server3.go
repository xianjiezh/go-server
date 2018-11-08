package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func handle1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handle1"))
}

func handle2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handle2"))
}

// 
func makeHandlers(handlers ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			handler(w, r)
		}
	}
}

func main() {
	router := httprouter.New()
	http.HandleFunc("/", makeHandlers(handle1, handle2))
	if err := http.ListenAndServe(":12345", router); err != nil {
		fmt.Println("start errors", err)
	}
}
