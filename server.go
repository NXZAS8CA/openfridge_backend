package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello world")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	log.Print("Starting server at port 8080\n")

	if err :=http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
