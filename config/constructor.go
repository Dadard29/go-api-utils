package config

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
)

func NewAccessor(path string, verbose bool) *Accessor {
	logger := log.NewLogger("accessor", logLevel.LevelFromBool(verbose))

	file, err := newFile(path)
	logger.CheckErrFatal(err)
	logger.Debug("Loaded config file")

	env := newEnv()

	return &Accessor{
		configFilePath: path,
		file:           file,
		env:            env,
		logger:         logger,
	}
}
