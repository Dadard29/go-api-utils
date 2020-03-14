package test

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/Dadard29/go-api-utils/test"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	name := "TestLogger"
	level := logLevel.DEBUG

	logger := log.NewLogger(name, level)

	expectedName := strings.ToUpper(name)
	actualName := logger.Name()

	if actualName != expectedName {
		test.AssertError(t, "Wrong format of logger name", expectedName, actualName)
	}

	actualLevel := logger.Level()

	if logger.Level() != level {
		test.AssertError(t, "Wrong level set", level, actualLevel)
	}
}

func TestLog(t *testing.T) {
	tables := []struct {
		levelInit     int
		levelLog      int
		levelReturned int
	}{
		{logLevel.DEBUG, logLevel.DEBUG, logLevel.DEBUG},
		{logLevel.INFO, logLevel.DEBUG, -1},
		{logLevel.INFO, logLevel.INFO, logLevel.INFO},
		{logLevel.WARNING, logLevel.DEBUG, -1},
		{logLevel.WARNING, logLevel.INFO, -1},
		{logLevel.WARNING, logLevel.WARNING, logLevel.WARNING},
		{logLevel.ERROR, logLevel.ERROR, logLevel.ERROR},
		{logLevel.FATAL, logLevel.FATAL, logLevel.FATAL},
	}

	for _, v := range tables {
		logger := log.NewLogger("testLogger", v.levelInit)
		levelReturned := logger.Log("testMessage", v.levelLog)
		if levelReturned != v.levelReturned {
			test.AssertError(t, "Wrong returned level", v.levelReturned, levelReturned)
		}
	}
}
