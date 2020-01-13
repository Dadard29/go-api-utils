package service

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"net/http"
)

var infosObj infos
var apiLogger log.Logger
var routeList map[string]func(w http.ResponseWriter, r *http.Request)

func addJsonHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func routesListRoute(w http.ResponseWriter, r *http.Request) {
	var routeNameList []string
	for k, _ := range routeList {
		routeNameList = append(routeNameList, k)
	}

	addJsonHeader(w)

	err := json.NewEncoder(w).Encode(routeNameList)
	apiLogger.CheckErr(err)
}

func infosRoute(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)
	err := json.NewEncoder(w).Encode(infosObj)
	apiLogger.CheckErr(err)
}

func healthRoute(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)
	err := json.NewEncoder(w).Encode(struct {
		Status bool
	}{
		Status: true,
	})
	apiLogger.CheckErr(err)
}

func NewService(routes map[string]func(w http.ResponseWriter, r *http.Request), serverConfig map[string]string,
	infosConfig map[string]string, verbose bool) Service {
		var err error
		infosObj, err = newInfos(infosConfig)
		if err != nil {
			panic(err)
		}

		apiLogger = log.NewLogger(infosObj.Title, logLevel.LevelFromBool(verbose))

		if _, check := routes["/infos"]; ! check {
			routes["/infos"] = infosRoute
		}

		if _, check := routes["/health"]; ! check {
			routes["/health"] = healthRoute
		}

		if _, check := routes["/routes"]; ! check {
			routes["/routes"] = routesListRoute
		}

		routeList = routes

		router := newRouter(routes)

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


