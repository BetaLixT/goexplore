package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	opr := flag.String("opr", "def", "operation")
	msg := flag.String("msg", "Hello", "message")
	flag.Parse()
	if *opr == "send" {
		Producer(*msg)
	} else if *opr == "recv" {
		Reciever()
	} else {
		fmt.Print("No valid argument\n")
		os.Exit(-1)
	}
}
