package async_logger

import (
	"runtime"
	"strings"
)

func getCaller() string {
	_, file, _, _ := runtime.Caller(1)
	idx := strings.LastIndexByte(file, '/')
	idx = strings.LastIndexByte(file[:idx], '/')
	return file[idx+1:]
}
