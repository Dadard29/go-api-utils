package config

import "github.com/Dadard29/go-api-utils/log"

type Accessor struct {
	configFilePath string
	file           *file
	env            *env
	logger         log.Logger
}

type file struct {
	config map[string]map[string]map[string]string
}

type env struct {
	config map[string]string
}
