package logger

import "errors"

type Level int

const (
	INFO Level = iota
	WARN
	ERROR
	FATAL
	DEBUG
	TRACE
)

var levels = map[string]Level{
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
	"debug": DEBUG,
	"trace": TRACE,
}

var representations = map[Level]string{
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

func getLevel(id string) (level Level, err error) {
	level = FATAL
	if val, ok := levels[id]; ok {
		level = val
		return
	}
	err = errors.New("Unknown logging level: " + id)
	return
}

func getLevelRep(level Level) string {
	return representations[level]
}
