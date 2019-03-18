package main

import (
	"github.com/az-art/tops/pkg/tops"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tops", tops.HandlerTops)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
