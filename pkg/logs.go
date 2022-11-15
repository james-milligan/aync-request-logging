package async_logger

type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
)

func (a *AsyncLogger) Debug(reqID string, msg string) {
	a.loggers[reqID].buffer <- messageWrapper{
		level: DEBUG,
		msg:   msg,
	}
}

func (a *AsyncLogger) Info(reqID string, msg string) {
	a.loggers[reqID].buffer <- messageWrapper{
		level: INFO,
		msg:   msg,
	}
}

func (a *AsyncLogger) Warn(reqID string, msg string) {
	a.loggers[reqID].buffer <- messageWrapper{
		level: WARNING,
		msg:   msg,
		file:  getCaller(),
	}
}

func (a *AsyncLogger) Error(reqID string, msg string) {
	a.loggers[reqID].buffer <- messageWrapper{
		level: ERROR,
		msg:   msg,
		file:  getCaller(),
	}
}
