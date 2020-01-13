package service

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"net/http"
)

var infosObj infos
var connectorObj *database.Connector
var apiLogger log.Logger
//var routeList map[string]func(w http.ResponseWriter, r *http.Request)

func infosRoute(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(infosObj)
	apiLogger.CheckErr(err)
}

func healthRoute(w http.ResponseWriter, r *http.Request) {
	dbStatus := true
	err := connectorObj.Orm.DB().Ping()
	if err != nil {
		dbStatus = false
	}

	err = json.NewEncoder(w).Encode(struct {
		status bool
		dbStatus bool
	}{
		status: true,
		dbStatus: dbStatus,
	})
	apiLogger.CheckErr(err)
}

func NewService(routes map[string]func(w http.ResponseWriter, r *http.Request), serverConfig map[string]string,
	infosConfig map[string]string, connector *database.Connector, verbose bool) Service {
		var err error
		infosObj, err = newInfos(infosConfig)
		if err != nil {
			panic(err)
		}

		apiLogger = log.NewLogger(infosObj.Title, logLevel.LevelFromBool(verbose))

		if _, check := routes["/infos"]; ! check {
			routes["/infos"] = infosRoute
		}

		if _, check := routes["/health"]; check {
			routes["/health"] = healthRoute
		}

		//routeList = routes
		router := newRouter(routes)

		connectorObj = connector

		server, err := newServer(router, serverConfig, apiLogger)
		if err != nil {
			panic(err)
		}

		return Service{
			srv:    server,
			infos:  infosObj,
			logger: apiLogger,
		}
}


