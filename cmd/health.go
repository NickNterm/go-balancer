package main

import (
	"log"
	"sync"
	"time"

	"github.com/NickNterm/go-balancer/internal/healtchecker"
)

func (app *application) HealthCheckEveryTime() {
	go func() {
		ticker := time.NewTicker(time.Duration(app.config.HealthTicker) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			var wg sync.WaitGroup

			for i := range app.proxies {
				wg.Add(1)
				go func(proxy *Proxy) {
					defer wg.Done()
					CheckProxy(proxy)
				}(&app.proxies[i])
			}
			wg.Wait()
		}
	}()
}

func calculateAverage(oldAvg int32, value int32) int32 {
	return int32(float64(oldAvg) + 0.5*(float64(value)-float64(oldAvg)))
}

func CheckProxy(proxy *Proxy) {
	log.Printf("url: %s avrg: %d health: %t", proxy.addr, proxy.avgResponse, proxy.isHealthy)
	success, time := healtchecker.CheckHealth(proxy.addr)
	proxy.isHealthy = success
	if success {
		proxy.avgResponse = calculateAverage(proxy.avgResponse, time)
	}
}
