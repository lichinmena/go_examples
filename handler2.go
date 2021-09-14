package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The timeis:" + tm))
}

func main() {
	mux := http.NewServeMux()

	th := http.HandlerFunc(timeHandler)
	mux.Handle("/time", th)

	log.Println("Listenin...")
	http.ListenAndServe(":3000", mux)
}
