package main

import "fmt"
import "net/http"

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello again"))
	})

	fmt.Println("server started at localhost:80")
	http.ListenAndServe(":8080", nil)
}
