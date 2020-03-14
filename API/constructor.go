package API

import (
	"github.com/Dadard29/go-api-utils/config"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/Dadard29/go-api-utils/service"
)

func NewAPI(
	name string,
	configPath string,
	routeList service.RouteMapping,
	verbose bool,
) API {
	var api API

	logger := log.NewLogger(name, logLevel.LevelFromBool(verbose))

	accessor := config.NewAccessor(configPath, verbose)

	serverConfig, err := accessor.GetSubcategoryFromFile("api", "server")
	logger.CheckErrFatal(err)
	infosConfig, err := accessor.GetSubcategoryFromFile("api", "infos")
	logger.CheckErrFatal(err)

	if _, check := infosConfig["version"]; !check {
		infosConfig["version"] = accessor.GetEnv("VERSION")
	}

	if _, check := serverConfig["corsOrigin"]; !check {
		corsOrigin := accessor.GetEnv("CORS_ORIGIN")
		if corsOrigin == "" {
			logger.Fatal("missing configuration value for CORS_ORIGIN")
		} else {
			serverConfig["corsOrigin"] = corsOrigin
		}
	}

	// the connector is an optional element of the API. The dev user is in charge to initialize it
	// Api.Databaase = database.NewConnector(...)
	// thus, the dev user is free to implement a route to check the health of the db connection
	api.Config = accessor
	api.Logger = logger
	api.Database = nil

	server := service.NewService(routeList, serverConfig, infosConfig, verbose)
	api.Service = server

	return api
}
