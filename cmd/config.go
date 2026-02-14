package main

import (
	"encoding/json"
	"log"
	"os"
)

type ServerConfig struct {
	Addr   string  `json:"addr"`
	Weight float32 `json:"weight"`
}

type Config struct {
	Algorithm    Algorithm      `json:"algorithm"`
	Addr         string         `json:"addr"`
	Servers      []ServerConfig `json:"servers"`
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
