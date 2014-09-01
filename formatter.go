package logger

import (
	"fmt"
	"strings"
)

type formatterI interface {
	Format(message string) string
}

type formatter struct {
	level   Level
	newLine bool
	layout  string
}

func (f formatter) format(message string, args []interface{}) string {

	output := f.layout

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	output = strings.Replace(output, "{level}", getLevelRep(f.level), -1)
	output = strings.Replace(output, "{message}", message, -1)

	if f.newLine {
		output += "\n"
	}

	return output
}

func buildInfoFormatter(settings settings) formatter {
	return formatter{
		level:  INFO,
		layout: settings.Layout,
	}
}

func buildInfolnFormatter(settings settings) formatter {
	return formatter{
		level:   INFO,
		newLine: true,
		layout:  settings.Layout,
	}
}

func buildInfofFormatter(settings settings) formatter {
	return formatter{
		level:   INFO,
		newLine: true,
		layout:  settings.Layout,
	}
}

func buildWarnFormatter(settings settings) formatter {
	return formatter{
		level:  WARN,
		layout: settings.Layout,
	}
}

func buildWarnlnFormatter(settings settings) formatter {
	return formatter{
		level:   WARN,
		newLine: true,
		layout:  settings.Layout,
	}
}

func buildWarnfFormatter(settings settings) formatter {
	return formatter{
		level:   WARN,
		newLine: true,
		layout:  settings.Layout,
	}
}
