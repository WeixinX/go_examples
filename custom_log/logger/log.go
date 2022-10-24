package logger

import (
	"fmt"
	"io"
	"log"
)

const (
	ERROR = iota
	WARN
	INFO
	DEBUG
)

type Log struct {
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Debug *log.Logger
	level int
}

func New(logLevel int, out io.Writer) *Log {
	l := &Log{
		Error: log.New(out, "[ERROR] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix),
		Warn:  log.New(out, "[WARN] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix),
		Info:  log.New(out, "[INFO] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix),
		Debug: log.New(out, "[DEBUG] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix),
		level: logLevel,
	}
	return l
}

func (l *Log) Errorf(format string, v ...interface{}) {
	if l.level < ERROR {
		return
	}
	// 若 calldepath = 1，则记录文件为 log.go；若 = 2，则为 Infoln 调用处，后面的类似
	err := l.Error.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		return
	}
}

func (l *Log) Warnf(format string, v ...interface{}) {
	if l.level < WARN {
		return
	}
	err := l.Warn.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		return
	}
}
func (l *Log) Infof(format string, v ...interface{}) {
	if l.level < INFO {
		return
	}
	err := l.Info.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		return
	}
}
func (l *Log) Debugf(format string, v ...interface{}) {
	if l.level < DEBUG {
		return
	}
	err := l.Debug.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		return
	}
}
