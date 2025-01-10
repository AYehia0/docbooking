package logger

import "log"

type Logger struct {
	l *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		l: log.New(log.Writer(), "", log.LstdFlags),
	}
}

// Info logs a message at the info level.
func (l *Logger) Info(msg string) {
	l.l.Printf("INFO: %s", msg)
}

// Infof logs a formatted message at the info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.l.Printf("INFO: "+format, v...)
}

// Error logs a message at the error level.
func (l *Logger) Error(msg string) {
	l.l.Printf("ERROR: %s", msg)
}

// Errorf logs a formatted message at the error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.l.Printf("ERROR: "+format, v...)
}

// Fatal logs a message at the fatal level.
func (l *Logger) Fatal(msg string) {
	l.l.Fatalf("FATAL: %s", msg)
}

// Fatalf logs a formatted message at the fatal level.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.l.Fatalf("FATAL: "+format, v...)
}
