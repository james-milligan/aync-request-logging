package async_logger

import (
	"fmt"

	"go.uber.org/zap"
)

type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
)

func (a *AsyncLogger) Debug(reqID string, msg string) {
	if a.loggers[reqID] != nil && a.loggers[reqID].buffer != nil {
		a.loggers[reqID].buffer <- messageWrapper{
			level: DEBUG,
			msg:   msg,
		}
	} else {
		a.parentLogger.Error(
			fmt.Sprintf("log attempted for closed channel %s", reqID),
			zap.String("requestID", reqID),
		)
	}
}

func (a *AsyncLogger) Info(reqID string, msg string) {
	if a.loggers[reqID] != nil && a.loggers[reqID].buffer != nil {
		a.loggers[reqID].buffer <- messageWrapper{
			level: INFO,
			msg:   msg,
		}
	} else {
		a.parentLogger.Error(
			fmt.Sprintf("log attempted for closed channel %s", reqID),
			zap.String("requestID", reqID),
		)
	}
}

func (a *AsyncLogger) Warn(reqID string, msg string) {
	if a.loggers[reqID] != nil && a.loggers[reqID].buffer != nil {
		a.loggers[reqID].buffer <- messageWrapper{
			level:  WARNING,
			msg:    msg,
			source: getCaller,
		}
	} else {
		a.parentLogger.Error(
			fmt.Sprintf("log attempted for closed channel %s", reqID),
			zap.String("requestID", reqID),
		)
	}
}

func (a *AsyncLogger) Error(reqID string, msg string) {
	if a.loggers[reqID] != nil && a.loggers[reqID].buffer != nil {
		a.loggers[reqID].buffer <- messageWrapper{
			level:  ERROR,
			msg:    msg,
			source: getCaller,
		}
	} else {
		a.parentLogger.Error(
			fmt.Sprintf("log attempted for closed channel %s", reqID),
			zap.String("requestID", reqID),
		)
	}
}
