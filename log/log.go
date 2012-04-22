package log

import (
	stdlog "log"
)

type Log struct{}

func (l *Log) Info(fmt string, args ...interface{}) {
	l.Printf("INFO ", fmt, args...)
}

func (l *Log) Error(fmt string, args ...interface{}) {
	l.Printf("ERROR ", fmt, args...)
}

func (l *Log) Printf(prefix, fmt string, args ...interface{}) {
	old := stdlog.Prefix()
	stdlog.SetPrefix(prefix)
	defer stdlog.SetPrefix(old)
	stdlog.Printf(fmt, args...)
}
