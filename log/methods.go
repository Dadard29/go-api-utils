package log

import (
	"fmt"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/fatih/color"
	"strings"
)

// if prefix and suffix of s is `, no format is applied
func (logger *Logger) Log(s string, level int) int {
	var content string
	if strings.HasPrefix(s, "`") && strings.HasSuffix(s, "`") {
		content = strings.Trim(s, "`")
	} else {
		content = strings.ToLower(s)
	}

	if level == logLevel.FATAL {
		panic(s)
	}

	if level >= logger.level {
		levelStr, err := getLevelName(level)
		if err != nil {
			loggerError(err.Error())
		}

		c := getLevelColor(level)

		message := fmt.Sprintf("%s %s", levelStr, content)

		color.Set(c)
		logger.logger.Println(message)
		color.Unset()

		return level
	} else {
		return -1
	}
}

func (logger *Logger) Debug(s string) {
	logger.Log(s, logLevel.DEBUG)
}

func (logger *Logger) Info(s string) {
	logger.Log(s, logLevel.INFO)
}

func (logger *Logger) Warning(s string) {
	logger.Log(s, logLevel.WARNING)
}

func (logger *Logger) Error(s string) {
	logger.Log(s, logLevel.ERROR)
}

func (logger *Logger) Fatal(s string) {
	logger.Log(s, logLevel.FATAL)
}
