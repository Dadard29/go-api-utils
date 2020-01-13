package log

import "log"

//Logger structure to manage logging
type Logger struct {
	logger *log.Logger
	name   string
	level  int
	levelStr string
}

func (logger *Logger) Logger() *log.Logger {
	return logger.logger
}

func (logger *Logger) LevelStr() string {
	return logger.levelStr
}

func (logger *Logger) Level() int {
	return logger.level
}

func (logger *Logger) Name() string {
	return logger.name
}

