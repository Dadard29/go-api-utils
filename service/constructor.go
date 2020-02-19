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

func addJsonHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func routesListRoute(w http.ResponseWriter, r *http.Request) {
	var routeNameList []string
	for k, _ := range routeList.Mapping {
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

func NewService(routes RouteMapping, serverConfig map[string]string,
	infosConfig map[string]string, verbose bool) Service {
		var err error
		infosObj, err = newInfos(infosConfig)
		if err != nil {
			panic(err)
		}

		apiLogger = log.NewLogger(infosObj.Title, logLevel.LevelFromBool(verbose))

		if _, check := routes.Mapping["/infos"]; ! check {
			routes.Mapping["/infos"] = Route{infosRoute, []string{http.MethodGet}}
		}

		if _, check := routes.Mapping["/health"]; ! check {
			routes.Mapping["/health"] = Route{healthRoute, []string{http.MethodGet}}
		}

		if _, check := routes.Mapping["/routes"]; ! check {
			routes.Mapping["/routes"] = Route{routesListRoute, []string{http.MethodGet}}
		}

		routeList = routes

		router := newRouter(routes)

		return Service{
			srv: nil,
			router: router,
			infos:  infosObj,
			logger: apiLogger,
			serverConfig: serverConfig,
		}
}


