package main

import (
	"log"
	"net/http"
)

type application struct {
	config  Config
	proxies []Proxy
}

func (app *application) run() error {
	handler := http.HandlerFunc(app.balancerHandler)
	server := http.Server{
		Addr:    app.config.Addr,
		Handler: handler,
	}

	log.Printf("Running in %s", server.Addr)

	return server.ListenAndServe()
}
