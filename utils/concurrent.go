package utils

import (
	"log"
	"runtime/debug"
)

// ReportPanic 异常上报
func ReportPanic(err interface{}) {
	log.Printf("panic: %v, stack: %s", err, debug.Stack())
}

// CatchPanic 异常捕获
func CatchPanic() {
	if err := recover(); err != nil {
		ReportPanic(err)
	}
}

// SafeGo 安全协程
func SafeGo(f func()) {
	go func() {
		defer CatchPanic()
		f()
	}()
}
