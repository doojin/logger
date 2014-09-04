package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatter_test_Before(t *testing.T) {
	defaultTiming = mockedTiming{}
}

type testcase struct {
	input          string
	formatter      formatter
	expectedOutput string
	args           []interface{}
}

var testcases = []testcase{
	// newLine: true
	testcase{
		input: "dummy",
		formatter: formatter{
			newLine: true,
			layout:  "{message}",
		},
		expectedOutput: "dummy\n",
	},
	// newLine: false
	testcase{
		input: "dummy",
		formatter: formatter{
			newLine: false,
			layout:  "{message}",
		},
		expectedOutput: "dummy",
	},
	// formatting
	testcase{
		input: "dummy %2.2f",
		formatter: formatter{
			layout: "{message}",
		},
		expectedOutput: "dummy 12.35",
		args:           []interface{}{12.3456},
	},
	// formatting and new line
	testcase{
		input: "dummy %v dummy",
		formatter: formatter{
			newLine: true,
			layout:  "{message}",
		},
		expectedOutput: "dummy cat dummy\n",
		args:           []interface{}{"cat"},
	},
	// level: INFO
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  INFO,
		},
		expectedOutput: "[INFO] dummy",
	},
	// level: WARN
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  WARN,
		},
		expectedOutput: "[WARN] dummy",
	},
	// level: ERROR
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  ERROR,
		},
		expectedOutput: "[ERROR] dummy",
	},
	// level: FATAL
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  FATAL,
		},
		expectedOutput: "[FATAL] dummy",
	},
	// level: DEBUG
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  DEBUG,
		},
		expectedOutput: "[DEBUG] dummy",
	},
	// level: TRACE
	testcase{
		input: "dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
			level:  TRACE,
		},
		expectedOutput: "[TRACE] dummy",
	},
	// default level
	testcase{
		input: "dummy %v dummy",
		formatter: formatter{
			layout: "[{level}] {message}",
		},
		expectedOutput: "[INFO] dummy 15 dummy",
		args:           []interface{}{15},
	},
	// formatting with multiple arguments
	testcase{
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
	testcase{
		input: "dummy",
		formatter: formatter{
			layout:     "[{time}] {message}",
			level:      ERROR,
			timeFormat: "Mon, 02 Jan 2006 15:04:05",
		},
		expectedOutput: "[Wed, 02 Jan 1985 03:04:05] dummy",
	},
}

func Test_format_ShouldFormatCorrectly(t *testing.T) {
	for i := range testcases {
		testcase := testcases[i]
		actualOutput := testcase.formatter.format(testcase.input, testcase.args...)
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

func Test_formatter_test_After(t *testing.T) {
	defaultTiming = defTimingOriginal
}
