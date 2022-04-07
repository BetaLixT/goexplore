package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	fmt.Println("vim-go")
	godotenv.Load()
	key := os.Getenv("InstrumentationKey")

	// - insights
	client := appinsights.NewTelemetryClient(key)

	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})

	// - zap
	lvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	ins := zapcore.WriteSyncer(NewInsightsSync(client))
	enc := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(enc, ins, lvl),
	)

	logger := zap.New(core)
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("Register Insights",
		// Structured context as loosely typed key-value pairs.
		"insightsKey", key,
	)

	// fmt.Printf("Utilizing key: %s\n", key)

	// client.TrackTrace("Log from test logger", contracts.Information)
	// client.Channel().Flush()

}

type InsightsSink struct {
	clnt appinsights.TelemetryClient
}

var _ zapcore.WriteSyncer = (*InsightsSink)(nil)

func NewInsightsSync(clnt appinsights.TelemetryClient) *InsightsSink {
	return &InsightsSink{
		clnt: clnt,
	}
}

func (ins *InsightsSink) Sync() error {
	fmt.Println("Waiting for insights to sync...")
	ins.clnt.Channel().Close()
	select {
	case <-ins.clnt.Channel().Close(10 * time.Second):
		// Ten second timeout for retries.

		// If we got here, then all telemetry was submitted
		// successfully, and we can proceed to exiting.
	case <-time.After(30 * time.Second):
		// Thirty second absolute timeout.  This covers any
		// previous telemetry submission that may not have
		// completed before Close was called.

		// There are a number of reasons we could have
		// reached here.  We gave it a go, but telemetry
		// submission failed somewhere.  Perhaps old events
		// were still retrying, or perhaps we're throttled.
		// Either way, we don't want to wait around for it
		// to complete, so let's just exit.
	}
	return nil
}

func (ins *InsightsSink) Write(data []byte) (int, error) {
	ins.clnt.TrackTrace(string(data), contracts.Information)
	return len(data), nil
}
