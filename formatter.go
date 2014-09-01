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

func (f formatter) format(message string, args []interface{}) string {

	output := f.layout

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	output = strings.Replace(output, "{level}", getLevelRep(f.level), -1)
	output = strings.Replace(output, "{message}", message, -1)

	currentTime := defTiming.getCurrentTime().Format(f.timeFormat)
	output = strings.Replace(output, "{time}", currentTime, -1)

	if f.newLine {
		output += "\n"
	}

	return output
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
