package logger

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_logger_test_Before(t *testing.T) {
	defaultTiming = mockedTiming{}
}

type formatMessageTestCase struct {
	levelId  string
	message  string
	args     []interface{}
	expected string
}

var formatMessageTestCases = []formatMessageTestCase{
	// Unknown level
	formatMessageTestCase{
		levelId:  "dummy",
		message:  "dummy",
		expected: "03:04:05 [INFO] Unknown logging level: dummy\n",
	},
	// [INFO] message
	formatMessageTestCase{
		levelId:  "info",
		message:  "dummy",
		expected: "03:04:05 [INFO] dummy\n",
	},
	// [WARN] message
	formatMessageTestCase{
		levelId:  "warn",
		message:  "dummy",
		expected: "03:04:05 [WARN] dummy\n",
	},
	// [INFO] message with arguments
	formatMessageTestCase{
		levelId:  "info",
		message:  "dummy %v",
		args:     []interface{}{"argument1"},
		expected: "03:04:05 [INFO] dummy argument1\n",
	},
	// [WARN] message with arguments
	formatMessageTestCase{
		levelId:  "warn",
		message:  "dummy %v",
		args:     []interface{}{"argument1"},
		expected: "03:04:05 [WARN] dummy argument1\n",
	},
}

func Test_formatMessage_ShouldFormatMessageCorrectly(t *testing.T) {
	l := New()
	for _, testCase := range formatMessageTestCases {
		actual := l.formatMessage(testCase.levelId, testCase.message, testCase.args...)
		assert.Equal(t, testCase.expected, actual)
	}
}

func Test_getWriter_ShouldGetConsoleWriterById(t *testing.T) {
	logger := New()
	writer := logger.getWriter("console")
	assert.Equal(t, os.Stdout, writer)
}

func Test_getWriter_ShouldGetDefaultWriterForUnknownId(t *testing.T) {
	logger := New()
	writer := logger.getWriter("dummy")
	assert.Equal(t, os.Stdout, writer)
}

func Test_logger_test_After(t *testing.T) {
	defaultTiming = defTimingOriginal
}
