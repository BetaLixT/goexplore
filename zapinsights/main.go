package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
)

func main() {
	fmt.Println("vim-go")
	godotenv.Load()
	key := os.Getenv("InstrumentationKey")
	fmt.Printf("Utilizing key: %s\n", key)
	client := appinsights.NewTelemetryClient(key)

	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})

	client.TrackTrace("Log from test logger", contracts.Information)
	client.Channel().Flush()
	fmt.Scanln()

}
