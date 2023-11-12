package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port = flag.Int("port", 3333, "The port that the server should listen on. Defaults to port 3333.")

func main() {
	flag.Parse()

	http.HandleFunc("/", index)

	log.Printf("Server listening on port %d", *port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}
