package main

import (
	"log"
	"net/http"
)

func (app *application) balancerHandler(w http.ResponseWriter, r *http.Request) {
	p, err := app.algorithm.ProcessRequest(app.proxies)
	if err != nil {
		return
	}
	log.Printf("SelectedProxy has %s, avg %d, health %t", p.Addr, p.AvgResponse, p.IsHealthy)
	p.Proxy.ServeHTTP(w, r)
}
