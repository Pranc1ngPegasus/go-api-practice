// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/configuration"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/controller/router"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/controller/server"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure"
	"github.com/google/wire"
)

func initialize() *http.Server {
	wire.Build(
		configuration.Get,
		infrastructure.NewRDBConnector,
		router.NewRouter,
		server.NewServer,
	)

	return nil
}
