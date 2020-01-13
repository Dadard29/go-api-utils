package API

import (
	"github.com/Dadard29/go-api-utils/config"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/service"
)

type API struct {
	Accessor *config.Accessor
	Logger log.Logger
	Connector *database.Connector
	Service service.Service
}