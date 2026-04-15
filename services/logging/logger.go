package logging

import (
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/fx/fxevent"
)

type logger struct {
}

func (*logger) LogDebug(format string, v ...interface{}) {
	slog.Debug(fmt.Sprintf(format, v...))
}

func (*logger) LogInfo(format string, v ...interface{}) {
	slog.Info(fmt.Sprintf(format, v...))
}

func (*logger) LogWarn(format string, v ...interface{}) {
	slog.Warn(fmt.Sprintf(format, v...))
}

func (*logger) LogError(format string, v ...interface{}) {
	slog.Error(fmt.Sprintf(format, v...))
}

func (*logger) LogFatal(format string, v ...interface{}) {
	slog.Error(fmt.Sprintf(format, v...))
	slog.Error(fmt.Sprintf("FATAL ERROR ENCOUNTERED, TERMINATING EXECUTION..."))
	os.Exit(1)
}

func (*logger) LogEvent(event fxevent.Event) {
	s := fxevent.SlogLogger{
		Logger: slog.Default(),
	}
	s.LogEvent(event)
}

func NewLogger() (LogWrapper, FxLogger) {
	l := &logger{}
	return l, l
}

type LogWrapper interface {
	LogDebug(format string, v ...interface{})
	LogInfo(format string, v ...interface{})
	LogWarn(format string, v ...interface{})
	LogError(format string, v ...interface{})
}

type FxLogger interface {
	LogEvent(fxevent.Event)
}
