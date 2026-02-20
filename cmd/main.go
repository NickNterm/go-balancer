package main

import (
	"log"

	"github.com/NickNterm/go-balancer/internal/algorithms"
	"github.com/NickNterm/go-balancer/internal/proxy"
)

func main() {
	c := GetConfig()
	proxies, err := proxy.CreateReverseProxies(c.Servers)

	if err != nil {
		log.Printf("Can't create reverse proxies %v", err)
	}

	var algorithm Algorithm
	switch c.Algorithm {
	case RoundRobin:
		algorithm = &algorithms.RoundRobin{
			CurrentIndex: 0,
		}
	case Random:
		algorithm = &algorithms.Random{}
	case LeastResponseTime:
		algorithm = &algorithms.LeastResponseTime{}
	}
	app := application{
		config:    c,
		proxies:   proxies,
		algorithm: algorithm,
	}

	app.HealthCheckEveryTime()

	app.run()
}
