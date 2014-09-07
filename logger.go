package logger

import (
	"io"
	"os"
)

// Logger log the information into console or a file
// Logger can be configgured by setting values to the properties of
// Logger.Settings
type Logger struct {
	// Configuration settings
	Settings settings
}

// Logger configuration
type settings struct {
	// The destination of logging ("default", "console", "file")
	Writer string
	// Format of messages, for example: "{time} [{level}] {message}"
	Layout string
	// Time format
	TimeFormat string
	// If Writer is "file", Filename should be provided. Specifies path
	// to the file where messages will be logged
	Filename string
}

// Returns a new instance of Logger
func New() *Logger {
	return &Logger{
		Settings: settings{
			Writer:     "default",
			Layout:     "{time} [{level}] {message}",
			TimeFormat: "15:04:05",
		},
	}
}

// Logs message to the writer
// levelId - Specifies the type of message ("info", "warn", "error",
// "fatal", "debug" "trace")
// message - message which will be formatted and logged
// args - arguments for formatting
func (l *Logger) Log(levelId string, message string, args ...interface{}) {
	formattedMessage := l.formatMessage(levelId, message, args...)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formattedMessage)
}

// Logs Info message without new line and message formatting
func (l *Logger) Info(message string) {
	formatter := buildInfoFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Info message with new line and without message formatting
func (l *Logger) Infoln(message string) {
	formatter := buildInfolnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Info message with new line and message formatting
func (l *Logger) Infof(message string, args ...interface{}) {
	formatter := buildInfofFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Logs Warn message without new line and message formatting
func (l *Logger) Warn(message string) {
	formatter := buildWarnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Warn message with new line and without message formatting
func (l *Logger) Warnln(message string) {
	formatter := buildWarnlnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Warn message with new line and message formatting
func (l *Logger) Warnf(message string, args ...interface{}) {
	formatter := buildWarnfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Logs Error message without new line and message formatting
func (l *Logger) Error(message string) {
	formatter := buildErrorFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Error message with new line and without message formatting
func (l *Logger) Errorln(message string) {
	formatter := buildErrorlnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Error message with new line and message formatting
func (l *Logger) Errorf(message string, args ...interface{}) {
	formatter := buildErrorfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Logs Fatal message without new line and message formatting
func (l *Logger) Fatal(message string) {
	formatter := buildFatalFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Fatal message with new line and without message formatting
func (l *Logger) Fatalln(message string) {
	formatter := buildFatallnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Fatal message with new line and message formatting
func (l *Logger) Fatalf(message string, args ...interface{}) {
	formatter := buildFatalfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Logs Debug message without new line and message formatting
func (l *Logger) Debug(message string) {
	formatter := buildDebugFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Debug message with new line and without message formatting
func (l *Logger) Debugln(message string) {
	formatter := buildDebuglnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Debug message with new line and message formatting
func (l *Logger) Debugf(message string, args ...interface{}) {
	formatter := buildDebugfFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Logs Trace message without new line and message formatting
func (l *Logger) Trace(message string) {
	formatter := buildTraceFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Trace message with new line and without message formatting
func (l *Logger) Traceln(message string) {
	formatter := buildTracelnFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message))
}

// Logs Trace message with new line and message formatting
func (l *Logger) Tracef(message string, args ...interface{}) {
	formatter := buildTracefFormatter(l.Settings)
	writer := l.getWriter(l.Settings.Writer)
	l.write(writer, formatter.format(message, args...))
}

// Formats message depending on levelId and formatting arguments
func (l *Logger) formatMessage(levelId string, message string, args ...interface{}) (result string) {
	// If level not supported by logger
	level, err := getLevel(levelId)
	if err != nil {
		result = buildErrorfFormatter(l.Settings).format("%v", err)
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

// Gets writer by it's name
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

// Gets writer for logging into file
func (l *Logger) getFileWriter() (file io.Writer) {
	file, _ = os.OpenFile(l.Settings.Filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	return
}

// Writes message to the writer
func (l *Logger) write(writer io.Writer, message string) {
	io.WriteString(writer, message)
	if l.Settings.Writer == "file" {
		writer.(*os.File).Close()
	}
}
