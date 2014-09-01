package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase struct {
	input          string
	formatter      formatter
	expectedOutput string
	args           []interface{}
}

var testcases = []testcase{
	{
		"dummy",
		formatter{
			newLine: true,
			layout:  "{message}",
		},
		"dummy\n",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			newLine: false,
			layout:  "{message}",
		},
		"dummy",
		[]interface{}{},
	},
	{
		"dummy %2.2f",
		formatter{
			layout: "{message}",
		},
		"dummy 12.35",
		[]interface{}{12.3456},
	},
	{
		"dummy %v dummy",
		formatter{
			newLine: true,
			layout:  "{message}",
		},
		"dummy cat dummy\n",
		[]interface{}{"cat"},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  INFO,
		},
		"[INFO] dummy",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  WARN,
		},
		"[WARN] dummy",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  ERROR,
		},
		"[ERROR] dummy",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  FATAL,
		},
		"[FATAL] dummy",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  DEBUG,
		},
		"[DEBUG] dummy",
		[]interface{}{},
	},
	{
		"dummy",
		formatter{
			layout: "[{level}] {message}",
			level:  TRACE,
		},
		"[TRACE] dummy",
		[]interface{}{},
	},
	{
		"dummy %v dummy",
		formatter{
			layout: "[{level}] {message}",
		},
		"[INFO] dummy 15 dummy",
		[]interface{}{15},
	},
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
    
    assert.Equal(t, "[INFO] dummy", actualOutput)
}

func Test_buildInfolnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
    logger := New()
    formatter := buildInfolnFormatter(logger.Settings)
    input := "dummy"
    
    actualOutput := formatter.format(input, []interface{}{})
    
    assert.Equal(t, "[INFO] dummy\n", actualOutput)
}

func Test_buildInfofFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
    logger := New()
    formatter := buildInfolnFormatter(logger.Settings)
    input := "dummy %v"
    
    actualOutput := formatter.format(input, []interface{}{"dummy"})
    
    assert.Equal(t, "[INFO] dummy dummy\n", actualOutput)
}

func Test_buildWarnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
    logger := New()
    formatter := buildWarnFormatter(logger.Settings)
    input := "dummy"
    
    actualOutput := formatter.format(input, []interface{}{})
    
    assert.Equal(t, "[WARN] dummy", actualOutput)
}

func Test_buildWarnlnFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
    logger := New()
    formatter := buildWarnlnFormatter(logger.Settings)
    input := "dummy"
    
    actualOutput := formatter.format(input, []interface{}{})
    
    assert.Equal(t, "[WARN] dummy\n", actualOutput)
}

func Test_buildWarnfFormatter_ShouldProvideCorrectFormatter(t *testing.T) {
    logger := New()
    formatter := buildWarnfFormatter(logger.Settings)
    input := "dummy %v"
    
    actualOutput := formatter.format(input, []interface{}{"dummy"})
    
    assert.Equal(t, "[WARN] dummy dummy\n", actualOutput)
}