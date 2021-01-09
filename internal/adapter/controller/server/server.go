package server

import (
	"fmt"
	"net/http"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/configuration"
	log "github.com/sirupsen/logrus"
)

func NewServer(handler http.Handler, config configuration.Config) *http.Server {
	port := config.ListenPort
	log.Printf("Listen at :%d", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	return server
}
