package logger

import (
	"fmt"
	"strings"
)

type formatterI interface {
	Format(message string) string
}

type formatter struct {
	level      Level
	newLine    bool
	layout     string
	timeFormat string
}

func (f formatter) format(message string, args ...interface{}) string {
	message = f.formatArgs(message, args...)
	output := f.formatLayout(message)
	output = f.formatNewLine(output)
	return output
}

func (f formatter) formatArgs(message string, args ...interface{}) string {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	return message
}

func (f formatter) formatLayout(message string) string {
	layout := f.layout
	layout = strings.Replace(layout, "{level}", getLevelRep(f.level), -1)
	layout = strings.Replace(layout, "{message}", message, -1)
	currentTime := defaultTiming.getCurrentTime().Format(f.timeFormat)
	layout = strings.Replace(layout, "{time}", currentTime, -1)
	return layout
}

func (f formatter) formatNewLine(message string) string {
	if f.newLine {
		message += "\n"
	}
	return message
}

func buildInfoFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildInfolnFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildInfofFormatter(settings settings) formatter {
	return formatter{
		level:      INFO,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildWarnFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildWarnlnFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildWarnfFormatter(settings settings) formatter {
	return formatter{
		level:      WARN,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildErrorFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildErrorlnFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}

func buildErrorfFormatter(settings settings) formatter {
	return formatter{
		level:      ERROR,
		newLine:    true,
		layout:     settings.Layout,
		timeFormat: settings.TimeFormat,
	}
}
