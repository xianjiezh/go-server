package main      
import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path == "/abc" {
		w.Write([]byte("abc"))
		return 
	}
	if r.URL.Path == "/xyz" {
		w.Write([]byte("xyz"))
		return
	}
	w.Write([]byte("index"))
}

func main()  {
	http.Handle("/", MyHandler{})
	if err := http.ListenAndServe(":12345", MyHandler{}); err != nil {
		fmt.Println("start http err", err)
	}
}
