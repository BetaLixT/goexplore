package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	protoparser "github.com/yoheimuta/go-protoparser/v4"
)

func run() int {
	flag.Parse()

	reader, err := os.Open("contracts.proto")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open contracts.proto, err %v\n", err)
		return 1
	}
	defer reader.Close()

	got, err := protoparser.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse, err %v\n", err)
		return 1
	}

	gotJSON, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal, err %v\n", err)
	}
	fmt.Print(string(gotJSON))
	return 0
}

func main() {
	os.Exit(run())
}
