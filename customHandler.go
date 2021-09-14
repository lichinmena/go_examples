package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	th1123 := &timeHandler{format: time.RFC1123} //Produce un puntero
	mux.Handle("/time/th1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339} //Produce un puntero
	mux.Handle("/time/th3339", th3339)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
