package tops

import (
	"fmt"
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

func HandlerTops(w http.ResponseWriter, req *http.Request) {
	var c Command
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
