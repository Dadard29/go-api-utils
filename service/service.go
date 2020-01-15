package service

import (
	"github.com/Dadard29/go-api-utils/log"
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
}

type Route struct {
	Handler func (w http.ResponseWriter, r *http.Request)
	Method string
}

type RouteMapping struct {
	Mapping map[string]Route
}


