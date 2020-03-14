package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

func (a *Service) Infos() map[string]string {
	return a.infos.toMap()
}

func (a *Service) InfosString() string {
	return a.infos.toString()
}

func (a *Service) CorsOrigin() string {
	return a.serverConfig["corsOrigin"]
}

func (a *Service) Start() {
	// a.srv.ListenAndServe()

	server, err := newServer(a.router, a.serverConfig, apiLogger)
	apiLogger.CheckErrFatal(err)

	a.srv = server

	a.logger.Info(fmt.Sprintf("starting the API service on %s", server.Addr))
	go func() {
		a.srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (a *Service) Stop() error {
	if a.srv == nil {
		msg := "server not started. start it first you idiot."
		a.logger.Error(msg)
		return errors.New(msg)
	}

	a.logger.Info("stopping the API service")
	var wait time.Duration
	ctx, _ := context.WithTimeout(context.Background(), wait)
	a.srv.Shutdown(ctx)

	a.srv = nil
	return nil
}

func (a *Service) IsRunning() bool {
	return a.srv != nil
}

func (a *Service) AddFileServer(prefix string, relativePath string) {
	fileServerRoot, err := filepath.Abs(relativePath)
	a.logger.CheckErr(err)

	slash := "/"
	if !strings.HasPrefix(prefix, slash) {
		prefix = slash + prefix
	}

	if !strings.HasSuffix(prefix, slash) {
		prefix = prefix + slash
	}

	a.router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir(fileServerRoot))))
}
