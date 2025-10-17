package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func format(level, message string) string {
	pc1, _, _, ok1 := runtime.Caller(3)
	pc2, _, _, ok2 := runtime.Caller(4)

	fn1, fn2 := "unknown", "unknown"
	if ok1 {
		fn := runtime.FuncForPC(pc1)
		parts := strings.Split(fn.Name(), "/")
		fn1 = parts[len(parts)-1]
	}
	if ok2 {
		fn := runtime.FuncForPC(pc2)
		parts := strings.Split(fn.Name(), "/")
		fn2 = parts[len(parts)-1]
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s] [%s -> %s] %s", timestamp, level, fn2, fn1, message)
}

func Infof(formatStr string, args ...interface{}) {
	log.Println(format("INFO", fmt.Sprintf(formatStr, args...)))
}

func Errorf(formatStr string, args ...interface{}) {
	log.Println(format("ERROR", fmt.Sprintf(formatStr, args...)))
}

func Fatalf(formatStr string, args ...interface{}) {
	log.Println(format("FATAL", fmt.Sprintf(formatStr, args...)))
	os.Exit(1)
}
