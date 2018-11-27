package errors

import (
	"io"
	"log"
)

// Log 收集錯誤訊息
type Log struct {
	callerLevel int
	writer      io.Writer
}

func (l *Log) Error(err string) {
	log.Printf("logerr: \n%s\n", err)
}
