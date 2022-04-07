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
	fmt.Scanln()

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
	fmt.Println("Syncing")
	ins.clnt.Channel().Flush()
	return nil
}

func (ins *InsightsSink) Write(data []byte) (int, error) {
	ins.clnt.TrackTrace(string(data), contracts.Information)
	return len(data), nil
}
