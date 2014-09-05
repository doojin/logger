package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatter_test_Before(t *testing.T) {
	defaultTiming = mockedTiming{}
}

type formatTestcase struct {
	input          string
	formatter      formatter
	expectedOutput string
	args           []interface{}
}

type formatLayoutTestcase struct {
	formatter      formatter
	message        string
	expectedOutput string
}

var formatTestcases = []formatTestcase{
	// newLine: true
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			newLine: true,
			layout:  "{message}",
		},
		expectedOutput: "dummy\n",
	},
	// newLine: false
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			newLine: false,
			layout:  "{message}",
		},
		expectedOutput: "dummy",
	},
	// formatting
	formatTestcase{
		input: "dummy %2.2f",
		formatter: formatter{
			layout: "{message}",
		},
		expectedOutput: "dummy 12.35",
		args:           []interface{}{12.3456},
	},
	// formatting and new line
	formatTestcase{
		input: "dummy %v dummy",
		formatter: formatter{
			newLine: true,
			layout:  "{message}",
		},
		expectedOutput: "dummy cat dummy\n",
		args:           []interface{}{"cat"},
	},
	// level: INFO
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  INFO,
		},
		expectedOutput: "[INFO] dummy",
	},
	// level: WARN
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  WARN,
		},
		expectedOutput: "[WARN] dummy",
	},
	// level: ERROR
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  ERROR,
		},
		expectedOutput: "[ERROR] dummy",
	},
	// level: FATAL
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  FATAL,
		},
		expectedOutput: "[FATAL] dummy",
	},
	// level: DEBUG
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  DEBUG,
		},
		expectedOutput: "[DEBUG] dummy",
	},
	// level: TRACE
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  TRACE,
		},
		expectedOutput: "[TRACE] dummy",
	},
	// default level
	formatTestcase{
		input: "dummy %v dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
		},
		expectedOutput: "[INFO] dummy 15 dummy",
		args:           []interface{}{15},
	},
	// formatting with multiple arguments
	formatTestcase{
		input: "dummy %v dummy %v dummy",
		formatter: formatter{
			newLine: true,
			layout:  "[{level}] {message}",
			level:   WARN,
		},
		expectedOutput: "[WARN] dummy black dummy white dummy\n",
		args:           []interface{}{"black", "white"},
	},
	// time formatting
	formatTestcase{
		input: "dummy",
		formatter: formatter{
			layout:     "[{time}] {message}",
			level:      ERROR,
			timeFormat: "Mon, 02 Jan 2006 15:04:05",
		},
		expectedOutput: "[Wed, 02 Jan 1985 03:04:05] dummy",
	},
}

var formatLayoutTestcases = []formatLayoutTestcase{
	// INFO
	formatLayoutTestcase{
		formatter: formatter{
			level:      INFO,
			layout:     "{level} {time} {message}",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "INFO 03:04:05 dummy",
	},
	// WARN
	formatLayoutTestcase{
		formatter: formatter{
			level:      WARN,
			layout:     "{level} {time} {message}",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "WARN 03:04:05 dummy",
	},
	// ERROR
	formatLayoutTestcase{
		formatter: formatter{
			level:      ERROR,
			layout:     "{level} {time} {message}",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "ERROR 03:04:05 dummy",
	},
	// DEBUG
	formatLayoutTestcase{
		formatter: formatter{
			level:      DEBUG,
			layout:     "{level} {time} {message}",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "DEBUG 03:04:05 dummy",
	},
	// TRACE
	formatLayoutTestcase{
		formatter: formatter{
			level:      TRACE,
			layout:     "{level} {time} {message}",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "TRACE 03:04:05 dummy",
	},
	// Empty layout
	formatLayoutTestcase{
		formatter: formatter{
			level:      WARN,
			layout:     "",
			timeFormat: "15:04:05",
		},
		message:        "dummy",
		expectedOutput: "",
	},
	// Custom time format
	formatLayoutTestcase{
		formatter: formatter{
			level:      WARN,
			layout:     "{time}",
			timeFormat: "Mon, 02 Jan 2006 15:04:05",
		},
		message:        "dummy",
		expectedOutput: "Wed, 02 Jan 1985 03:04:05",
	},
}

func Test_format_ShouldFormatCorrectly(t *testing.T) {
	for i := range formatTestcases {
		testcase := formatTestcases[i]
		actualOutput := testcase.formatter.format(testcase.input, testcase.args...)
		assert.Equal(t, testcase.expectedOutput, actualOutput)
	}
}

func Test_formatLayout_ShouldFormatCorrectly(t *testing.T) {
	for i := range formatLayoutTestcases {
		testcase := formatLayoutTestcases[i]
		actualOutput := testcase.formatter.formatLayout(testcase.message)
		assert.Equal(t, testcase.expectedOutput, actualOutput)
	}
}

func Test_buildInfoFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfoFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [INFO] dummy", actualOutput)
}

func Test_buildInfolnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfolnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [INFO] dummy\n", actualOutput)
}

func Test_buildInfofFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfolnFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [INFO] dummy dummy\n", actualOutput)
}

func Test_buildWarnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [WARN] dummy", actualOutput)
}

func Test_buildWarnlnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnlnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [WARN] dummy\n", actualOutput)
}

func Test_buildWarnfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnfFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [WARN] dummy dummy\n", actualOutput)
}

func Test_buildErrorFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildErrorFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [ERROR] dummy", actualOutput)
}

func Test_buildErrorlnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildErrorlnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [ERROR] dummy\n", actualOutput)
}

func Test_buildErrorfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildErrorfFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [ERROR] dummy dummy\n", actualOutput)
}

func Test_buildFatalFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildFatalFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [FATAL] dummy", actualOutput)
}

func Test_buildFatallnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildFatallnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [FATAL] dummy\n", actualOutput)
}

func Test_buildFatalfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildFatalfFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [FATAL] dummy dummy\n", actualOutput)
}

func Test_buildDebugFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildDebugFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [DEBUG] dummy", actualOutput)
}

func Test_buildDebuglnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildDebuglnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [DEBUG] dummy\n", actualOutput)
}

func Test_buildDebugfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildDebugfFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [DEBUG] dummy dummy\n", actualOutput)
}

func Test_buildTraceFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildTraceFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [TRACE] dummy", actualOutput)
}

func Test_buildTracelnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildTracelnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input)

	assert.Equal(t, "03:04:05 [TRACE] dummy\n", actualOutput)
}

func Test_buildTracefFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildTracefFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"}...)

	assert.Equal(t, "03:04:05 [TRACE] dummy dummy\n", actualOutput)
}

func Test_formatArgs_ShouldFormatArgumentsIfTheyExist(t *testing.T) {
	f := formatter{}
	input := "dummy %v %v"

	actualOutput := f.formatArgs(input, "arg1", "arg2")

	assert.Equal(t, "dummy arg1 arg2", actualOutput)
}

func Test_formatArgs_ShouldNotFormatArgumentsIfTheyDontExist(t *testing.T) {
	f := formatter{}
	input := "dummy %v %v"

	actualOutput := f.formatArgs(input)

	assert.Equal(t, "dummy %v %v", actualOutput)
}

func Test_formatNewLine_ShouldAddNewLineIfNeed(t *testing.T) {
	f := formatter{
		newLine: true,
	}
	input := "dummy"

	actualOutput := f.formatNewLine(input)

	assert.Equal(t, "dummy\n", actualOutput)
}

func Test_formatNewLine_ShouldFNotAddNewLineIfNotNeed(t *testing.T) {
	f := formatter{
		newLine: false,
	}
	input := "dummy"

	actualOutput := f.formatNewLine(input)

	assert.Equal(t, "dummy", actualOutput)
}

func Test_formatter_test_After(t *testing.T) {
	defaultTiming = defTimingOriginal
}
