package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Command struct {
	com  string
	args []string
	w    http.ResponseWriter
	res  result
}

type result struct {
	out []byte
	err error
}

func main() {
	var c Command
	http.HandleFunc("/top", c.executeTop)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (c *Command) executeTop(w http.ResponseWriter, req *http.Request) {
	c.com = "top"
	c.args = []string{"-b", "-n", "1"}
	c.w = w
	c.exec()
}

func (c *Command) exec() {
	cmd := exec.Command(c.com, c.args...)
	if c.res.out, c.res.err = cmd.Output(); c.res.err != nil {
		println(c.res.err.Error())
	}
	fmt.Fprint(c.w, string(c.res.out))
}

