package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 3810

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
