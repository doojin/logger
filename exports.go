package logger

var defaultLogger = New()
var Settings = &defaultLogger.Settings

func Log(levelId string, message string, args ...interface{}) {
	defaultLogger.Log(levelId, message, args...)
}

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

func Fatal(message string) {
	defaultLogger.Fatal(message)
}

func Fatalln(message string) {
	defaultLogger.Fatalln(message)
}

func Fatalf(message string, args ...interface{}) {
	defaultLogger.Fatalf(message, args...)
}

func Debug(message string) {
	defaultLogger.Debug(message)
}

func Debugln(message string) {
	defaultLogger.Debugln(message)
}

func Debugf(message string, args ...interface{}) {
	defaultLogger.Debugf(message, args...)
}

func Trace(message string) {
	defaultLogger.Trace(message)
}

func Traceln(message string) {
	defaultLogger.Traceln(message)
}

func Tracef(message string, args ...interface{}) {
	defaultLogger.Tracef(message, args...)
}
