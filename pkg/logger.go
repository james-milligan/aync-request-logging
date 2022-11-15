package async_logger

import (
	"context"

	"github.com/rs/xid"
	"go.uber.org/zap"
)

type AsyncLogger struct {
	parentLogger *zap.Logger
}

type requestLogger struct {
	logger *zap.Logger
	buffer chan (messageWrapper)
}

type messageWrapper struct {
	msg   string
	level int
}

func (a *AsyncLogger) StartLogger(ctx context.Context) (context.CancelFunc, string) {
	ctx, cancel := context.WithCancel(ctx)
	reqID := xid.New().String()
	go a.registerLogger(ctx, reqID)
	return cancel, reqID
}

func (a AsyncLogger) registerLogger(ctx context.Context, reqID string) {

}
