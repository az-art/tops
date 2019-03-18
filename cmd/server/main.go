package main

import (
	"fmt"
	"github.com/az-art/tops/pkg/tops"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	var c Command
	http.HandleFunc("/top", c.executeTop)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
