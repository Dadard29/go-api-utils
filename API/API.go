package API

import (
	"github.com/Dadard29/go-api-utils/config"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/service"
)

type API struct {
	Config   *config.Accessor
	Logger   log.Logger
	Database *database.Connector
	Service  service.Service
}

type response struct {
	Status  bool
	Message string
	Content interface{}
}
