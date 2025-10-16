package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	colorInfo  = "\033[32m"
	colorError = "\033[31m"
	colorDebug = "\033[36m"
	colorFatal = "\033[35m"
	colorReset = "\033[0m"
)

func logWithLevel(level string, format string, args ...interface{}) {
	pc1, _, _, ok1 := runtime.Caller(2)
	pc2, _, _, ok2 := runtime.Caller(3)

	func1, func2 := "???", "???"

	if ok1 {
		if fn := runtime.FuncForPC(pc1); fn != nil {
			func1 = shortFuncName(fn.Name())
		}
	}

	if ok2 {
		if fn := runtime.FuncForPC(pc2); fn != nil {
			func2 = shortFuncName(fn.Name())
		}
	}

	color := colorReset
	switch level {
	case "INFO":
		color = colorInfo
	case "ERROR":
		color = colorError
	case "DEBUG":
		color = colorDebug
	case "FATAL":
		color = colorFatal
	}

	msg := fmt.Sprintf(format, args...)
	log.Printf("%s[%s] [%s] [%s] -> %s%s", color, level, func2, func1, msg, colorReset)

	if level == "FATAL" {
		os.Exit(1)
	}
}

func shortFuncName(full string) string {
	if idx := strings.LastIndex(full, "."); idx != -1 {
		return full[idx+1:]
	}
	return filepath.Base(full)
}

func Infof(format string, args ...interface{})  {
	logWithLevel("INFO", format, args...)
}

func Errorf(format string, args ...interface{}) {
	logWithLevel("ERROR", format, args...)
}

func Debugf(format string, args ...interface{}) {
	logWithLevel("DEBUG", format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logWithLevel("FATAL", format, args...)
}
