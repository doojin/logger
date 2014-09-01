package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var defTimingOriginal = defTiming

func Test_Before(t *testing.T) {
	defTiming = mockedTiming{}
}

type testcase struct {
	input          string
	formatter      formatter
	expectedOutput string
	args           []interface{}
}

var testcases = []testcase{
	// newLine: true
	{
		"dummy",
		formatter{
			newLine: true,
			layout:  "{message}",
		},
		"dummy\n",
		[]interface{}{},
	},
	// newLine: false
	{
		"dummy",
		formatter{
			newLine: false,
			layout:  "{message}",
		},
		"dummy",
		[]interface{}{},
	},
	// formatting
	{
		"dummy %2.2f",
		formatter{
			layout: "{message}",
		},
		"dummy 12.35",
		[]interface{}{12.3456},
	},
	// formatting and new line
	{
		"dummy %v dummy",
		formatter{
			newLine: true,
			layout:  "{message}",
		},
		"dummy cat dummy\n",
		[]interface{}{"cat"},
	},
	// level: INFO
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  INFO,
		},
		"[INFO] dummy",
		[]interface{}{},
	},
	// level: WARN
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  WARN,
		},
		"[WARN] dummy",
		[]interface{}{},
	},
	// level: ERROR
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  ERROR,
		},
		"[ERROR] dummy",
		[]interface{}{},
	},
	// level: FATAL
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  FATAL,
		},
		"[FATAL] dummy",
		[]interface{}{},
	},
	// level: DEBUG
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  DEBUG,
		},
		"[DEBUG] dummy",
		[]interface{}{},
	},
	// level: TRACE
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  TRACE,
		},
		"[TRACE] dummy",
		[]interface{}{},
	},
	// default level
	{
		"dummy %v dummy",
		formatter{
			layout: "[{level}] {message}",
		},
		"[INFO] dummy 15 dummy",
		[]interface{}{15},
	},
	// formatting with multiple arguments
	{
		"dummy %v dummy %v dummy",
		formatter{
			newLine: true,
			layout:  "[{level}] {message}",
			level:   WARN,
		},
		"[WARN] dummy black dummy white dummy\n",
		[]interface{}{"black", "white"},
	},
	// time formatting
	{
		"dummy",
		formatter{
			layout:     "[{time}] {message}",
			level:      ERROR,
			timeFormat: "Mon, 02 Jan 2006 15:04:05",
		},
		"[Wed, 02 Jan 1985 03:04:05] dummy",
		[]interface{}{},
	},
}

func Test_format_ShouldFormatCorrectly(t *testing.T) {
	for i := range testcases {
		testcase := testcases[i]
		actualOutput := testcase.formatter.format(testcase.input, testcase.args)
		assert.Equal(t, testcase.expectedOutput, actualOutput)
	}
}

func Test_buildInfoFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfoFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input, []interface{}{})

	assert.Equal(t, "03:04:05 [INFO] dummy", actualOutput)
}

func Test_buildInfolnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfolnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input, []interface{}{})

	assert.Equal(t, "03:04:05 [INFO] dummy\n", actualOutput)
}

func Test_buildInfofFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildInfolnFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"})

	assert.Equal(t, "03:04:05 [INFO] dummy dummy\n", actualOutput)
}

func Test_buildWarnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input, []interface{}{})

	assert.Equal(t, "03:04:05 [WARN] dummy", actualOutput)
}

func Test_buildWarnlnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnlnFormatter(logger.Settings)
	input := "dummy"

	actualOutput := formatter.format(input, []interface{}{})

	assert.Equal(t, "03:04:05 [WARN] dummy\n", actualOutput)
}

func Test_buildWarnfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
	logger := New()
	formatter := buildWarnfFormatter(logger.Settings)
	input := "dummy %v"

	actualOutput := formatter.format(input, []interface{}{"dummy"})

	assert.Equal(t, "03:04:05 [WARN] dummy dummy\n", actualOutput)
}

func Test_After(t *testing.T) {
	defTiming = defTimingOriginal
}
