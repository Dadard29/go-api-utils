package service

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var title string

func (i infos) toString() string {
	return fmt.Sprintf("\n===== * =====\n" +
		"%s\n" +
		"version: \t%s\n" +
		"description: \t%s\n" +
		"contact email: \t%s\n" +
		"===== * =====\n",
		i.Title, i.Version, i.Description, i.ContactEmail)
}

func (i infos) toMap() map[string]string {
	return map[string]string {
		"Title": i.Title,
		"Version": i.Version,
		"Description": i.Description,
		"ContactEmail": i.ContactEmail,
		"License": i.License,
		"LicenseUrl": i.LicenseUrl,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s: %s %s %s", strings.ToUpper(title), r.Host, r.Method, r.URL)
		fmt.Println(msg)

		next.ServeHTTP(w, r)
	})
}

func newRouter(routeMapping RouteMapping) *mux.Router {
	router := mux.NewRouter()

	for i, h := range routeMapping.Mapping {
		router.HandleFunc(i, h.Handler).Methods(h.Method...)
	}

	router.Use(loggingMiddleware)
	return router
}

func newInfos(config map[string]string) (infos, error) {

	title = config["title"]

	version := config["version"]
	description := config["description"]
	contactEmail := config["contactEmail"]
	license := config["license"]
	licenseUrl := config["licenseUrl"]

	if strings.Contains(title, " ") {
		return infos{}, errors.New("wrong service title format: no space allowed")
	}

	if version == "" {
		return infos{}, errors.New("no version given")
	}

	if ! strings.Contains(contactEmail, "@") {
		return infos{}, errors.New("wrong email format: missing '@'")
	}

	if ! strings.Contains(licenseUrl, "http://") && ! strings.Contains(licenseUrl, "https://") {
		return infos{}, errors.New("wrong license url format: unknown scheme")
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
	writeTimeout, err := strconv.Atoi(serverConfig["writeTimeout"])
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