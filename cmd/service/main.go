package main

import (
	"net/http"

	"github.com/matias-inc/muxapi/middleware"
	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/helpers"
	"github.com/rodrigopmatias/daddy-api/resources"
)

var (
	logger = helpers.GetLogger()
	config = helpers.GetConfig()
)

func main() {
	handle := router.NewRouterHandler()

	handle.Use(&middleware.TookMiddleware{})
	handle.Use(middleware.NewAccessLoggerMiddleware(logger))
	handle.Use(&middleware.RequestIdMiddleware{})

	base := router.NewRouterGroup("/v1")
	resources.InitRouter(base)
	handle.IncludeRouterGroup(base)

	service := http.Server{
		Addr:    config.AppAddr,
		Handler: handle,
	}

	logger.Infof("Server is running on http://%s", config.AppAddr)
	if err := service.ListenAndServe(); err != nil {
		logger.Err(err.Error())
	}
}
