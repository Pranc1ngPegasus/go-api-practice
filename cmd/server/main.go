package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/configuration"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	configuration.Load()

	server := initialize()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("server did not boot. err:%v", err)
		return
	}
}
