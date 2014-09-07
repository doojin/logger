package logger

import (
	"fmt"
	"strings"
)

// Formatter formats message according to settings set by user
type formatter struct {
	level      Level
	newLine    bool
	layout     string
	timeFormat string
}

// Formats message
func (f formatter) format(message string, args ...interface{}) string {
	message = f.formatArgs(message, args...)
	output := f.formatLayout(message)
	output = f.formatNewLine(output)
	return output
}

// Formats arguments of message, placing them into defined placeholders
func (f formatter) formatArgs(message string, args ...interface{}) string {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	return message
}

// Formats layout (specifies position of message level, message text and time
// in the output)
func (f formatter) formatLayout(message string) string {
	layout := f.layout
	layout = strings.Replace(layout, "{level}", getLevelRep(f.level), -1)
	layout = strings.Replace(layout, "{message}", message, -1)
	currentTime := defaultTiming.getCurrentTime().Format(f.timeFormat)
	layout = strings.Replace(layout, "{time}", currentTime, -1)
	return layout
}

// Adds new line in the end of message if required
func (f formatter) formatNewLine(message string) string {
	if f.newLine {
		message += "\n"
	}
	return message
}

// Builds formatter for Info messages (without new line)
func buildInfoFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Info messages (with new line)
func buildInfolnFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Info messages (with new line)
func buildInfofFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Warning messages (without new line)
func buildWarnFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Warning messages (with new line)
func buildWarnlnFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Warning messages (with new line)
func buildWarnfFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Error messages (without new line)
func buildErrorFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Error messages (with new line)
func buildErrorlnFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Error messages (with new line)
func buildErrorfFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Fatal messages (without new line)
func buildFatalFormatter(settings settings) formatter {
	return formatter{
		level:      FATAL,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Fatal messages (with new line)
func buildFatallnFormatter(settings settings) formatter {
	return formatter{
		level:      FATAL,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Fatal messages (with new line)
func buildFatalfFormatter(settings settings) formatter {
	return formatter{
		level:      FATAL,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Debug messages (without new line)
func buildDebugFormatter(settings settings) formatter {
	return formatter{
		level:      DEBUG,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Debug messages (with new line)
func buildDebuglnFormatter(settings settings) formatter {
	return formatter{
		level:      DEBUG,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Debug messages (with new line)
func buildDebugfFormatter(settings settings) formatter {
	return formatter{
		level:      DEBUG,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Trace messages (without new line)
func buildTraceFormatter(settings settings) formatter {
	return formatter{
		level:      TRACE,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Trace messages (with new line)
func buildTracelnFormatter(settings settings) formatter {
	return formatter{
		level:      TRACE,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

// Builds formatter for Trace messages (with new line)
func buildTracefFormatter(settings settings) formatter {
	return formatter{
		level:      TRACE,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}
