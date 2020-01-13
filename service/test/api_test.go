package test

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/service"
	"github.com/Dadard29/go-api-utils/config"
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

var serverConfig, _ = accessor.GetSubcategoryFromFile("service", "server")
var infosConfig, _ = accessor.GetSubcategoryFromFile("service", "infos")

var routes = map[string]func(w http.ResponseWriter, r *http.Request) {
	"/test": testRoute,
}

var a = service.NewService(routes, serverConfig, infosConfig, true)

func TestRun(t *testing.T) {
	a.Start()
	defer a.Stop()
}
