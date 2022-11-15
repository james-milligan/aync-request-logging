package main

import (
	"context"
	"time"

	async_logger "github.com/james-milligan/aync-request-logging/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
		DisableCaller: true,
	}
	logger, _ := cfg.Build()
	x := async_logger.NewAsyncLogger(logger)

	cancel, reqID := x.StartLogger(context.Background())
	x.Info(reqID, "first message")
	x.Debug(reqID, "first message")
	x.Warn(reqID, "first message")
	x.Error(reqID, "first message")

	cancel()
	time.Sleep(10 * time.Second)
}
