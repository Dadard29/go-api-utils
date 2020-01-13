package service

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func (a *Service) Infos() map[string]string {
	return a.infos.toMap()
}

func (a *Service) InfosString() string {
	return a.infos.toString()
}

func (a *Service) Start() {
	a.logger.Info("starting the API service")
	// a.srv.ListenAndServe()
	go func() {
		a.srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (a *Service) Stop() {
	a.logger.Info("stopping the API service")
	var wait time.Duration
	ctx, _ := context.WithTimeout(context.Background(), wait)
	a.srv.Shutdown(ctx)
}

