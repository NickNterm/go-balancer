package main

import (
	"log"
	"math/rand"
	"net/http"
)

func (app *application) balancerHandler(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(len(app.config.Servers))
	proxy := app.proxies[index]
	log.Printf("SelectedProxy has %s, avg %d, health %t", proxy.addr, proxy.avgResponse, proxy.isHealthy)
	proxy.proxy.ServeHTTP(w, r)
}
