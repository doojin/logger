package logger

import (
	"io"
	"os"
)

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

func (l *Logger) Warnf(message string, args ...interface{}) {
	formatter := buildWarnfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) Error(message string) {
	formatter := buildErrorFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Errorln(message string) {
	formatter := buildErrorlnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	formatter := buildErrorfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) Fatal(message string) {
	formatter := buildFatalFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Fatalln(message string) {
	formatter := buildFatallnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	formatter := buildFatalfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) Debug(message string) {
	formatter := buildDebugFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Debugln(message string) {
	formatter := buildDebuglnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	formatter := buildDebugfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

func (l *Logger) Trace(message string) {
	formatter := buildTraceFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Traceln(message string) {
	formatter := buildTracelnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

func (l *Logger) Tracef(message string, args ...interface{}) {
	formatter := buildTracefFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
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
	case ERROR:
		result = buildErrorfFormatter(l.Settings).format(message, args...)
	case FATAL:
		result = buildFatalfFormatter(l.Settings).format(message, args...)
	case DEBUG:
		result = buildDebugfFormatter(l.Settings).format(message, args...)
	case TRACE:
		result = buildTracefFormatter(l.Settings).format(message, args...)
	default:
		result = buildInfofFormatter(l.Settings).format(message, args...)
	}
	return
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
