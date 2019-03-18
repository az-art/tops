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
	fileName := "webrequests.log"

	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to webrequests.log
	log.SetOutput(logFile)
	mux := http.NewServeMux()
	mux.HandleFunc("/tops", tops.HandlerTops)
	//http.HandleFunc("/tops", tops.HandlerTops)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
