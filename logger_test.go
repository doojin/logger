package logger

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

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
