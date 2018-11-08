package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func handle1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("handle1"))
}

func handle2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("handle2"))
}

func main() {
	router := httprouter.New()
	router.GET("/handle1", handle1)
	router.GET("/handle2", handle2)
	if err := http.ListenAndServe(":12345", router); err != nil {
		fmt.Println("start errors", err)
	}
}
