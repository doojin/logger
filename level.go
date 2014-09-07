package logger

import "errors"

type Level int

// Level constants supported by Logger
//
// INFO: Any information that can be interesting for application users, for
// example system lifecycle events
//
// WARN: Potentially harmful situation
//
// ERROR: Unexpected conditions, error events that might still allow the
// application to continue running.
//
// FATAL: Unexpected conditions which may abort application
//
// DEBUG: Diagnostically helpful information
//
// TRACE: A very detailed information about the event
const (
	INFO Level = iota
	WARN
	ERROR
	FATAL
	DEBUG
	TRACE
)

// Map of events, with id equal to event's id (alias)
var levels = map[string]Level{
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
	"debug": DEBUG,
	"trace": TRACE,
}

// Textual representation of each level
var representations = map[Level]string{
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

// Getting level by it's id (alias)
func getLevel(id string) (level Level, err error) {
	level = FATAL
	if val, ok := levels[id]; ok {
		level = val
		return
	}
	err = errors.New("Unknown logging level: " + id)
	return
}

// Getting textual representation of level
func getLevelRep(level Level) string {
	return representations[level]
}
