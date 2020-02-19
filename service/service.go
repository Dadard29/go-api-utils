package service

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/gorilla/mux"
	"net/http"
)

type infos struct {
	Title 	string
	Version string
	Description string
	ContactEmail string
	License string
	LicenseUrl string
}

type Service struct {
	srv *http.Server
	infos infos
	logger log.Logger
	router *mux.Router
	serverConfig map[string]string
}

func (a *Service) Router() *mux.Router {
	return a.router
}

type Route struct {
	Handler func (w http.ResponseWriter, r *http.Request)
	Method []string
}

type RouteMapping struct {
	Mapping map[string]Route
}


