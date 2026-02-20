package main

import "github.com/NickNterm/go-balancer/internal/proxy"

type AlgorithmType string

const (
	Random            AlgorithmType = "random"
	RoundRobin        AlgorithmType = "round-robin"
	LeastResponseTime AlgorithmType = "least-response-time"
)

type Algorithm interface {
	ProcessRequest(proxies []proxy.Proxy) (*proxy.Proxy, error)
}
