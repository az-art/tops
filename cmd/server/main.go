package main

import (
	"github.com/az-art/tops/pkg/tops"
	"log"
	"net/http"
	"os"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func main() {
	log.SetOutput(os.Stdout)
	mux := http.NewServeMux()
	mux.HandleFunc("/tops", tops.HandlerTops)
	log.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", RequestLogger(mux)))
}
