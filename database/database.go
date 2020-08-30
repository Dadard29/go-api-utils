package database

import (
	"github.com/Dadard29/go-api-utils/log"
	"gorm.io/gorm"
)

type DbConfig struct {
	usernameKey  string
	passwordKey  string
	databaseName string
	port         string
	host         string
}

// supports only sql type database
type Connector struct {
	Orm      *gorm.DB
	dbConfig DbConfig
	logger   log.Logger
}
