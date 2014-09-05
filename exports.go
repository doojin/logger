package logger

var defaultLogger = New()
var Settings = &defaultLogger.Settings

func Info(message string) {
	defaultLogger.Info(message)
}

func Infoln(message string) {
	defaultLogger.Infoln(message)
}

func Infof(message string, args ...interface{}) {
	defaultLogger.Infof(message, args...)
}

func Warn(message string) {
	defaultLogger.Warn(message)
}

func Warnln(message string) {
	defaultLogger.Warnln(message)
}

func Warnf(message string, args ...interface{}) {
	defaultLogger.Warnf(message, args...)
}

func Error(message string) {
	defaultLogger.Error(message)
}

func Errorln(message string) {
	defaultLogger.Errorln(message)
}

func Errorf(message string, args ...interface{}) {
	defaultLogger.Errorf(message, args...)
}
