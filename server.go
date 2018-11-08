package main      
import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LaLaLaLa"))
	})
	http.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abccccccc"))
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		fmt.Println("start http errors", err)
	}
}
