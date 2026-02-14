package main

import (
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	proxy     *httputil.ReverseProxy
	addr      string
	isHealthy bool

	// This will be a EMA
	avgResponse int32
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func CreateReverseProxies(servers []ServerConfig) ([]Proxy, error) {
	var proxies []Proxy

	for i := range servers {
		url, err := url.Parse(servers[i].Addr)
		if err != nil {
			return nil, err
		}
		proxy := Proxy{
			proxy:       NewProxy(url),
			addr:        servers[i].Addr,
			isHealthy:   true,
			avgResponse: 1,
		}
		proxies = append(proxies, proxy)
	}
	return proxies, nil
}
