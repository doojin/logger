package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase1 struct {
	key           string
	expectedLevel Level
	expectedError error
}

var testcases1 = []testcase1{
	{"info", INFO, nil},
	{"warn", WARN, nil},
	{"error", ERROR, nil},
	{"fatal", FATAL, nil},
	{"debug", DEBUG, nil},
	{"trace", TRACE, nil},
}

func Test_getLevel_shouldReturnCorrectLoggingLevel(t *testing.T) {
	for i := range testcases1 {
		c := testcases1[i]
		level, err := getLevel(c.key)
		assert.Equal(t, c.expectedLevel, level)
		assert.Equal(t, c.expectedError, err)
	}
}

func Test_getLevel_shouldReturnErrorForUnknownLoggingLevelId(t *testing.T) {
	level, err := getLevel("dummy")
	assert.Equal(t, FATAL, level)
	assert.Equal(t, "Unknown logging level: dummy", err.Error())
}
