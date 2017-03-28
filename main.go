package main

import (
	"os"

	"github.com/MartinSahlen/cloud-goa/app"
	"github.com/MartinSahlen/go-cloud-fn/shim"
	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
)

func main() {
	service := goa.New("adder")

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	service.WithLogger(goakit.New(logger))

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	adder := NewOperandsController(service)
	app.MountOperandsController(service, adder)

	shim.ServeHTTP(service.Mux.ServeHTTP)
}
