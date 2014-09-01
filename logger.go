package logger

import (
	"io"
	"os"
)

type LoggerI interface {
	Info(message string)
	Infoln(message string)
	Infof(message string, args ...interface{})
	Log(level Level, message string, args ...interface{})
}

type Logger struct {
	// Configuration settings
	Settings settings
	// Collection of writers
	Writers map[string]io.Writer
}

type settings struct {
	Destination string
	Layout      string
}

func New() *Logger {
	return &Logger{
		Settings: settings{
			Destination: "default",
			Layout:      "[{level}] {message}",
		},

		Writers: map[string]io.Writer{
			"default": os.Stdout,
			"console": os.Stdout,
		},
	}
}

func (l *Logger) Info(message string) {
	formatter := buildInfoFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, []interface{}{}))
}

func (l *Logger) Infoln(message string) {
	formatter := buildInfolnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, []interface{}{}))
}

func (l *Logger) Infof(message string, args ...interface{}) {
	formatter := buildInfofFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, args))
}

func (l *Logger) Warn(message string) {
	formatter := buildWarnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, []interface{}{}))
}

func (l *Logger) Warnln(message string) {
	formatter := buildWarnlnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, []interface{}{}))
}

func (l *Logger) Warnf(message string, args ...interface{}) {
	formatter := buildWarnfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Destination)
	io.WriteString(writer, formatter.format(message, args))
}

func (l *Logger) getWriter(key string) io.Writer {
	if writer, ok := l.Writers[key]; ok {
		return writer
	}
	return l.Writers["default"]
}
