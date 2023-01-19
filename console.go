package logger

import "log"

type consoleSink struct {
}

func (s *consoleSink) Send(format string, v ...any) {
	log.Printf(format, v...)
}

func NewConsoleSink() LogSink {
	return &consoleSink{}
}
