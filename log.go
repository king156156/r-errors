package errors

import (
	"io"
	"log"
)

type Log struct {
	callerLevel int
	writer      io.Writer
}

func (l *Log) Error(err string) {
	log.Printf("logerr: \n%s\n", err)
}
