package service

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

type infos struct {
	Title        string
	Version      string
	Description  string
	ContactEmail string
	License      string
	LicenseUrl   *url.URL
}

type Service struct {
	srv          *http.Server
	infos        infos
	logger       log.Logger
	router       *mux.Router
	serverConfig map[string]string
}

func (a *Service) Router() *mux.Router {
	return a.router
}

type Handler func(w http.ResponseWriter, r *http.Request)
type MethodMapping map[string]Handler
type RouteMapping map[string]Route

type Route struct {
	Description   string
	MethodMapping MethodMapping
}
