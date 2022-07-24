package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func main () { 

	// - zap
	lvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	ws, errws, err := openSinks()
	if err != nil {
	  panic(err)
	}

	ins := zapcore.WriteSyncer(&zapcore.BufferedWriteSyncer{
	  WS: ws,
	})
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(enc, ins, lvl),
	)

	logger := zap.New(
	  core,
	  zap.ErrorOutput(errws),
	  zap.AddStacktrace(zapcore.ErrorLevel),
	  zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			var samplerOpts []zapcore.SamplerOption
			return zapcore.NewSamplerWithOptions(
				core,
				time.Second,
				100,
				100,
				samplerOpts...,
			)
		}),
	)
	defer logger.Sync() // flushes buffer, if any
	logger.Info("test")
	logger.Error("error!!")
}

func openSinks() (zapcore.WriteSyncer, zapcore.WriteSyncer, error) {
	sink, closeOut, err := zap.Open("stderr")
	if err != nil {
		return nil, nil, err
	}
	errSink, _, err := zap.Open("stderr")
	if err != nil {
		closeOut()
		return nil, nil, err
	}
	return sink, errSink, nil
}
