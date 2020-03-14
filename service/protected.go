package service

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var title string
var corsOrigin string

func (i infos) toString() string {
	return fmt.Sprintf("\n===== * =====\n"+
		"%s\n"+
		"version: \t%s\n"+
		"description: \t%s\n"+
		"contact email: \t%s\n"+
		"===== * =====\n",
		i.Title, i.Version, i.Description, i.ContactEmail)
}

func (i infos) toMap() map[string]string {
	return map[string]string{
		"Title":        i.Title,
		"Version":      i.Version,
		"Description":  i.Description,
		"ContactEmail": i.ContactEmail,
		"License":      i.License,
		"LicenseUrl":   i.LicenseUrl.String(),
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s: %s %s %s", strings.ToUpper(title), r.Host, r.Method, r.URL)
		fmt.Println(msg)

		next.ServeHTTP(w, r)
	})
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", corsOrigin)
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.WriteHeader(http.StatusNoContent)
}

func newRouter(routeMapping RouteMapping, origin string) (*mux.Router, error) {
	router := mux.NewRouter()
	corsOrigin = origin

	for route, h := range routeMapping {
		for method, h := range h.MethodMapping {
			if method == http.MethodOptions {
				return nil, errors.New("OPTIONS method is reserved for CORS stuff")
			}
			router.HandleFunc(route, h).Methods(method)
		}

		router.HandleFunc(route, corsHandler).Methods(http.MethodOptions)
	}

	router.Use(loggingMiddleware)
	return router, nil
}

func newInfos(config map[string]string) (infos, error) {

	title = config["title"]

	version := config["version"]
	description := config["description"]
	contactEmail := config["contactEmail"]
	license := config["license"]
	licenseUrl, err := url.Parse(config["licenseUrl"])
	if err != nil {
		return infos{}, errors.New("error while parsing license url")
	}

	if strings.Contains(title, " ") {
		return infos{}, errors.New("wrong service title format: no space allowed")
	}

	if version == "" {
		return infos{}, errors.New("no version given")
	}

	if !strings.Contains(contactEmail, "@") {
		return infos{}, errors.New("wrong email format: missing '@'")
	}

	return infos{
		Title:        title,
		Version:      version,
		Description:  description,
		ContactEmail: contactEmail,
		License:      license,
		LicenseUrl:   licenseUrl,
	}, nil
}

func newServer(router *mux.Router, serverConfig map[string]string, logger log.Logger) (*http.Server, error) {
	addr := fmt.Sprintf(
		"%s:%s",
		serverConfig["host"],
		serverConfig["port"],
	)

	readTimeout, err := strconv.Atoi(serverConfig["readTimeout"])
	if err != nil {
		return nil, err
	}

	writeTimeout, err := strconv.Atoi(serverConfig["writeTimeout"])
	if err != nil {
		return nil, err
	}

	idleTimeout, err := strconv.Atoi(serverConfig["idleTimeout"])

	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:              addr,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       time.Duration(readTimeout) * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Duration(writeTimeout) * time.Second,
		IdleTimeout:       time.Duration(idleTimeout) * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          logger.Logger(),
		BaseContext:       nil,
		ConnContext:       nil,
	}, nil
}
