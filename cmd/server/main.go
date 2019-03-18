package main

import (
	"context"
	"flag"
	"github.com/az-art/tops/pkg/tops"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = flag.String("port", "8000", "listener port")

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)

	mux := http.NewServeMux()
	mux.HandleFunc("/tops", tops.HandlerTops)

	srv := &http.Server{
		Handler:      RequestLogger(mux),
		Addr:         ":" + *port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting Server. Listening on port " + *port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	//log.Fatal(srv.ListenAndServe())
	//log.Fatal(http.ListenAndServe("localhost:8000", RequestLogger(mux)))

	waitForShutdown(srv)
}

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

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
