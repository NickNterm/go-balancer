package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/NickNterm/go-balancer/internal/proxy"
)

type Config struct {
	Algorithm    AlgorithmType  `json:"algorithm"`
	Addr         string         `json:"addr"`
	Webhook      string         `json:"webhook"`
	Servers      []proxy.Server `json:"servers"`
	HealthTicker int            `json:"healthCheckDelay"`
}

func GetConfig() Config {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Printf("error %v", err)
	}
	c := Config{}
	err = json.Unmarshal(file, &c)

	if err != nil {
		log.Printf("Error parsing json %v", err)
	}

	return c
}
