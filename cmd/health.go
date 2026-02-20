package main

import (
	"log"
	"sync"
	"time"

	"github.com/NickNterm/go-balancer/internal/healtchecker"
	"github.com/NickNterm/go-balancer/internal/proxy"
)

func (app *application) HealthCheckEveryTime() {
	go func() {
		ticker := time.NewTicker(time.Duration(app.config.HealthTicker) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			var wg sync.WaitGroup

			for i := range app.proxies {
				wg.Add(1)
				go func(proxy *proxy.Proxy) {
					defer wg.Done()
					wasHealthy := proxy.IsHealthy
					CheckProxy(proxy)
					if wasHealthy != proxy.IsHealthy {
						app.InformServerStatus(*proxy)
					}
				}(&app.proxies[i])
			}
			wg.Wait()
		}
	}()
}

func calculateAverage(oldAvg int32, value int32) int32 {
	return int32(float64(oldAvg) + 0.5*(float64(value)-float64(oldAvg)))
}

func CheckProxy(p *proxy.Proxy) {
	log.Printf("url: %s avrg: %d health: %t", p.Addr, p.AvgResponse, p.IsHealthy)
	success, time := healtchecker.CheckHealth(p.Addr)
	p.IsHealthy = success
	if success {
		p.AvgResponse = calculateAverage(p.AvgResponse, time)
	}
}
