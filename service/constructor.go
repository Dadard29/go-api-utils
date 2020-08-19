package service

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"net/http"
)

var infosObj infos
var apiLogger log.Logger
var routeList RouteMapping

var jsonIndent = "    "

func addJsonHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", corsOrigin)
}

func routesListRoute(w http.ResponseWriter, r *http.Request) {
	var routeNameList []string
	for k := range routeList {
		routeNameList = append(routeNameList, k)
	}

	addJsonHeader(w)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)
	err := encoder.Encode(routeNameList)
	//err := API.BuildJsonResponse(true, "route list retrieved", routeNameList, w)
	apiLogger.CheckErr(err)
}

func infosRoute(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)
	err := encoder.Encode(infosObj.toMap())
	//err := API.BuildJsonResponse(true, "infos retrieved", infosObj.toMap(), w)
	apiLogger.CheckErr(err)
}

func healthRoute(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)
	err := encoder.Encode(struct {
		Status bool
	}{
		Status: true,
	})

	//err := API.BuildJsonResponse(true, "health check done", "", w)
	apiLogger.CheckErr(err)
}

func NewService(routes RouteMapping, serverConfig map[string]string,
	infosConfig map[string]string, verbose bool) Service {
	var err error
	infosObj, err = newInfos(infosConfig)
	if err != nil {
		panic(err)
	}

	apiLogger = log.NewLogger(infosObj.Title, logLevel.LevelFromBool(verbose))

	if _, check := routes["/infos"]; !check {
		routes["/infos"] = Route{"API infos", MethodMapping{
			http.MethodGet: infosRoute,
		}}
	}

	if _, check := routes["/health"]; !check {
		routes["/health"] = Route{"API health", MethodMapping{
			http.MethodGet: healthRoute,
		}}
	}

	if _, check := routes["/routes"]; !check {
		routes["/routes"] = Route{"route list", MethodMapping{
			http.MethodGet: routesListRoute,
		}}
	}

	routeList = routes

	router, err := newRouter(routes, serverConfig["corsOrigin"])
	apiLogger.CheckErrFatal(err)

	return Service{
		srv:          nil,
		router:       router,
		infos:        infosObj,
		logger:       apiLogger,
		serverConfig: serverConfig,
	}
}
