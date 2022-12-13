package main

import (
	"fmt"

	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("https://api.vivalakiara.com/adventsse/events?stream=advent")

	client.SubscribeRaw(func(msg *sse.Event) {
		// Got some data!
		fmt.Println(string(msg.Data))
	})
}
