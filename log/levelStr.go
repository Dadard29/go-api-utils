package log

import (
	"errors"
	"github.com/fatih/color"
)

const (
	debugStr   = "DEBUG"
	infoStr    = "INFO"
	warningStr = "WARNING"
	errorStr   = "ERROR"
	fatalStr   = "FATAL"
)

const (
	debugColor = color.FgWhite
	infoColor = color.FgGreen
	warningColor = color.FgYellow
	errorColor = color.FgRed
	fatalColor = color.FgHiRed
)

func getLevelColor(level int) color.Attribute {
	switch level {
	case 0:
		return debugColor
	case 1:
		return infoColor
	case 2:
		return warningColor
	case 3:
		return errorColor
	case 4:
		return fatalColor
	}

	return color.FgWhite
}

func getLevelName(level int) (string, error) {
	levelStr := ""
	switch level {
	case 0:
		levelStr = debugStr
	case 1:
		levelStr = infoStr
	case 2:
		levelStr = warningStr
	case 3:
		levelStr = errorStr
	case 4:
		levelStr = fatalStr
	}

	if levelStr == "" {
		msg := "unrecognized Logger level"
		loggerError(msg)
		return levelStr, errors.New(msg)
	}

	return levelStr, nil
}
