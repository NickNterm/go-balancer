package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/NickNterm/go-balancer/internal/proxy"
)

type StatusBody struct {
	ServerUrl string `json:"serverUrl"`
	Status    bool   `json:"status"`
}

func (app *application) InformServerStatus(p proxy.Proxy) error {
	bodyStruct := StatusBody{
		ServerUrl: p.Addr,
		Status:    p.IsHealthy,
	}

	bodyJson, err := json.Marshal(bodyStruct)

	if err != nil {
		return err
	}
	_, err = http.Post(app.config.Webhook, "application/json", bytes.NewBuffer(bodyJson))

	return err
}
