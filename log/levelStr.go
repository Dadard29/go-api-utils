package log

import "errors"

const (
	debugStr = "DEBUG"
	infoStr = "INFO"
	warningStr = "WARNING"
	errorStr = "ERROR"
	fatalStr = "FATAL"
)

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