package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	authHandler := httpauth.SimpleBasicAuth("luis", "123456789")
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", authHandler(finalHandler))

	log.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
