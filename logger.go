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
}

type settings struct {
	Writer     string
	Layout     string
	TimeFormat string
	Filename   string
}

func New() *Logger {
	return &Logger{
		Settings: settings{
			Writer:     "default",
			Layout:     "{time} [{level}] {message}",
			TimeFormat: "15:04:05",
		},
	}
}

func (l *Logger) Log(levelId string, message string, args ...interface{}) {
	formattedMessage := l.formatMessage(levelId, message, args...)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formattedMessage)
}

func (l *Logger) Info(message string) {
	formatter := buildInfoFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Infoln(message string) {
	formatter := buildInfolnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Infof(message string, args ...interface{}) {
	formatter := buildInfofFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) Warn(message string) {
	formatter := buildWarnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Warnln(message string) {
	formatter := buildWarnlnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) formatMessage(levelId string, message string, args ...interface{}) (result string) {
	// If level not supported by logger
	level, err := getLevel(levelId)
	if err != nil {
		result = buildInfofFormatter(l.Settings).format("%v", err)
		return
	}
	switch level {
	case INFO:
		result = buildInfofFormatter(l.Settings).format(message, args...)
	case WARN:
		result = buildWarnfFormatter(l.Settings).format(message, args...)
	default:
		result = buildInfofFormatter(l.Settings).format(message, args...)
	}
	return
}

func (l *Logger) Warnf(message string, args ...interface{}) {
	formatter := buildWarnfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) getWriter(writerName string) io.Writer {
	switch writerName {
	case "console":
		return os.Stdout
	case "file":
		return l.getFileWriter()
	default:
		return os.Stdout
	}
}

func (l *Logger) getFileWriter() (file io.Writer) {
	file, _ = os.OpenFile(l.Settings.Filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	return
}

func (l *Logger) write(writer io.Writer, message string) {
	io.WriteString(writer, message)
	if l.Settings.Writer == "file" {
		writer.(*os.File).Close()
	}
}
