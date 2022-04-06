package main

import (
	"fmt"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
)

func main() {
	fmt.Println("vim-go")
	client := appinsights.NewTelemetryClient("e6ec8366-a80a-46d8-9b32-567f8da4b6f1")

	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})

	client.TrackTrace("Log from test logger", contracts.Information)
	client.Channel().Flush()
	fmt.Scanln()

}
