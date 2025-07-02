package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world!"))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
