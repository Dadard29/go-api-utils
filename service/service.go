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


