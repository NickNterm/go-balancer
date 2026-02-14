package main

import (
	"log"
)

func main() {
	c := GetConfig()
	proxies, err := CreateReverseProxies(c.Servers)

	if err != nil {
		log.Printf("Can't create reverse proxies %v", err)
	}

	app := application{
		config:  c,
		proxies: proxies,
	}

	app.HealthCheckEveryTime()

	app.run()
}
