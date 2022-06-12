package main

import (
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello net/http!\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(helloGoHandler))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
