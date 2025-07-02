package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world!"))
	})
	http.HandleFunc("GET /health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok, available!"))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
