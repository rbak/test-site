package log

import (
    "fmt"
    "os"

    "github.com/inconshreveable/log15"
)

type Logger struct {
    log    log15.Logger
}

func (l *Logger) Trace(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Debug(message)
}

func (l *Logger) Debug(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Debug(message)
}

func (l *Logger) Info(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Info(message)
}

func (l *Logger) Warn(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Warn(message)
}

func (l *Logger) Error(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Error(message)
}

func (l *Logger) Critical(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Crit(message)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    l.log.Crit(message)
    Close()
    os.Exit(1)
}

func formatLogging(format string, v ...interface{}) (string) {
    var message string
    if len(v) > 0 {
        message = fmt.Sprintf(format, v)
    } else {
        message = format
    }
    return message
}
