package logger

// Instance of default logger
var defaultLogger = New()

// Pointer to settings of default logger
var Settings = &defaultLogger.Settings

// Logs custom level message
func Log(levelId string, message string, args ...interface{}) {
	defaultLogger.Log(levelId, message, args...)
}

// Logs Info message
func Info(message string) {
	defaultLogger.Info(message)
}

// Logs Info message with new line in the end
func Infoln(message string) {
	defaultLogger.Infoln(message)
}

// Logs formatted Info message with new line in the end
func Infof(message string, args ...interface{}) {
	defaultLogger.Infof(message, args...)
}

// Logs Warning message
func Warn(message string) {
	defaultLogger.Warn(message)
}

// Logs Warning message with new line in the end
func Warnln(message string) {
	defaultLogger.Warnln(message)
}

// Logs formatted Warning message with new line in the end
func Warnf(message string, args ...interface{}) {
	defaultLogger.Warnf(message, args...)
}

// Logs Error message
func Error(message string) {
	defaultLogger.Error(message)
}

// Logs Error message with new line in the end
func Errorln(message string) {
	defaultLogger.Errorln(message)
}

// Logs formatted Error message with new line in the end
func Errorf(message string, args ...interface{}) {
	defaultLogger.Errorf(message, args...)
}

// Logs Fatal message
func Fatal(message string) {
	defaultLogger.Fatal(message)
}

// Logs Fatal message with new line in the end
func Fatalln(message string) {
	defaultLogger.Fatalln(message)
}

// Logs formatted Fatal message with new line in the end
func Fatalf(message string, args ...interface{}) {
	defaultLogger.Fatalf(message, args...)
}

// Logs Debug message
func Debug(message string) {
	defaultLogger.Debug(message)
}

// Logs Fatal message with new line in the end
func Debugln(message string) {
	defaultLogger.Debugln(message)
}

// Logs formatted Fatal message with new line in the end
func Debugf(message string, args ...interface{}) {
	defaultLogger.Debugf(message, args...)
}

// Logs Trace message
func Trace(message string) {
	defaultLogger.Trace(message)
}

// Logs Trace message with new line in the end
func Traceln(message string) {
	defaultLogger.Traceln(message)
}

// Logs formatted Trace message with new line in the end
func Tracef(message string, args ...interface{}) {
	defaultLogger.Tracef(message, args...)
}
