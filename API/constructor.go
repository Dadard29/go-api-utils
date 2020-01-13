package API

import (
	"github.com/Dadard29/go-api-utils/config"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/Dadard29/go-api-utils/service"
	"net/http"
)

func NewAPI(
	name string,
	configPath string,
	modelsList []interface{},
	routeList map[string]func(w http.ResponseWriter, r *http.Request),
	verbose bool,
	) API {

		logger := log.NewLogger(name, logLevel.LevelFromBool(verbose))

		accessor := config.NewAccessor(configPath, verbose)

		dbConfig, err := accessor.GetSubcategoryFromFile("api", "db")
		logger.CheckErr(err)
		db := database.NewConnector(dbConfig, verbose, modelsList)

		serverConfig, err := accessor.GetSubcategoryFromFile("api", "server")
		logger.CheckErr(err)
		infosConfig, err := accessor.GetSubcategoryFromFile("api", "infos")
		server := service.NewService(routeList, serverConfig, infosConfig, db, verbose)

		return API{
			Accessor:  accessor,
			Logger: logger,
			Connector: db,
			Service: server,
		}
}

