package errors

import (
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// RError 定義方法
type RError interface {
	Write(msg string) RError
	Error() string
	ErrorPath() string
	Err() RError
	Clear()
	Log() RError
}

// Error 錯誤訊息
type Error struct {
	path, s, msg string
	callerLevel  int
	writer       io.Writer
}

// New 新增Error實體
func New(msg string) RError {
	return &Error{s: msg, writer: os.Stdout, callerLevel: 7}
}

// Error 印出
func (e *Error) Error() string {
	return e.s + " " + e.msg
}

// ErrorPath 印出
func (e *Error) ErrorPath() string {
	return e.Error() + e.path
}

// Err 紀錄路徑
func (e *Error) Err() RError {
	e.path = e.path + generateCallerList(e.callerLevel)
	return e
}

// Write 組合訊息
func (e *Error) Write(msg string) RError {
	e.path = e.path + generateCallerList(e.callerLevel)
	e.msg = e.msg + msg
	return e
}

// Clear 清除
func (e *Error) Clear() {
	e.path, e.msg = "", ""
}

func generateCallerList(callerLevel int) string {
	var callers strings.Builder

	for i := startCallerLevel; ; i++ {
		_, file, line, ok := runtime.Caller(i)

		if !ok || i == getCallerLevels(callerLevel) {
			break
		}

		var caller strings.Builder
		caller.WriteString("\n" + file)
		caller.WriteString(":")
		caller.WriteString(strconv.Itoa(line))

		callers.WriteString(caller.String())
		// callers.WriteString("\n")
	}

	return callers.String()
}

func getCallerLevels(callerLevel int) int {
	if callerLevel == -1 {
		return callerLevel
	}
	return startCallerLevel + (callerLevel - 1)
}

// Log 印log
func (e *Error) Log() RError {
	mylog.Error(e.Error())
	return e
}
