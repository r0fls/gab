package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//"strings"
)

func ab(args ...string) string {
	cmd := exec.Command("ab", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func main() {
	out := ab("-n", "10", "-c", "1", "http://127.0.0.1:8000/")
	fmt.Printf(out)
}
