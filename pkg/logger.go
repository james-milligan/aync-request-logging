package async_logger

import (
	"context"

	"github.com/rs/xid"
	"go.uber.org/zap"
)

type AsyncLogger struct {
	parentLogger *zap.Logger
	loggers      map[string]*requestLogger
}

type requestLogger struct {
	logger *zap.Logger
	buffer chan (messageWrapper)
}

type messageWrapper struct {
	msg   string
	level logLevel
	file  string
}

func NewAsyncLogger(parentLogger *zap.Logger) AsyncLogger {
	return AsyncLogger{
		loggers:      map[string]*requestLogger{},
		parentLogger: parentLogger,
	}

}

func (a *AsyncLogger) StartLogger(ctx context.Context, fields ...zap.Field) (context.CancelFunc, string) {
	ctx, cancel := context.WithCancel(ctx)
	reqID := xid.New().String()

	// create reference to requestLogger
	a.loggers[reqID] = &requestLogger{
		buffer: make(chan messageWrapper),
	}

	// start logging channel
	go a.log(ctx, fields, reqID)

	// unblock parent process
	return cancel, reqID
}

func (a *AsyncLogger) log(ctx context.Context, fields []zap.Field, reqID string) {
	if a.loggers[reqID] == nil {
		a.parentLogger.Error("received log from closed channel, reqID has been removed from pool")
	}
	// generate logger for request (the buffer may currently be filling)
	// for flagd this function would take the request and construct the fields async
	logger := a.parentLogger.With(fields...)
	a.loggers[reqID].logger = logger

	for {
		select {
		case message := <-a.loggers[reqID].buffer:
			logMessage(a.loggers[reqID].logger, message)
		case <-ctx.Done():
			// need to flush any remaining messages in the buffer
			for {
				select {
				case <-a.loggers[reqID].buffer:
				default:
					delete(a.loggers, reqID)
					return
				}
			}
		}
	}

}

func logMessage(logger *zap.Logger, message messageWrapper) {
	switch message.level {
	case DEBUG:
		logger.Debug(message.msg)
	case INFO:
		logger.Info(message.msg)
	case WARNING:
		logger.Warn(message.msg, zap.String("caller", message.file))
	case ERROR:
		logger.Error(message.msg, zap.String("caller", message.file))
	}
}
