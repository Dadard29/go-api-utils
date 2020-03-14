package test

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/config"
	"github.com/Dadard29/go-api-utils/service"
	"net/http"
	"testing"
)

func testRoute(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("salut")
	if err != nil {
		panic(err)
	}
}

var accessor = config.NewAccessor("service/test/config.json", false)

var serverConfig, _ = accessor.GetSubcategoryFromFile("api", "server")
var infosConfig, _ = accessor.GetSubcategoryFromFile("api", "infos")

var routes = service.RouteMapping{Mapping: map[string]service.Route{
	"/test": service.Route{Handler: testRoute, Method: []string{http.MethodGet}},
}}

var a = service.NewService(routes, serverConfig, infosConfig, true)

func TestRun(t *testing.T) {
	a.Start()
	defer a.Stop()
}
