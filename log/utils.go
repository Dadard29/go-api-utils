package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func loggerError(msg string) {
	fmt.Printf("LOGGER EXCEPTION: %s\n", msg)
}

func NewLogger(name string, level int) Logger {
	flags := log.Ldate | log.Ltime

	levelStr, err := getLevelName(level)
	if err != nil {
		panic(err)
	}

	nameUpper := strings.ToUpper(name)
	prefix := fmt.Sprintf("%s: ", nameUpper)

	logOutput := log.New(os.Stdout, prefix, flags)

	return Logger{
		logger:   logOutput,
		name:     nameUpper,
		level:    level,
		levelStr: levelStr,
	}
}
