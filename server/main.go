package main

import (
	"io"
	"log"
	"net/http"
)

const port = ":3333"

func main() {
	http.HandleFunc("/", index)

	log.Printf("Server listening on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}
