package proxy

import (
	"net/http/httputil"
	"net/url"
)

type Server struct {
	Addr   string  `json:"addr"`
	Weight float32 `json:"weight"`
}

type Proxy struct {
	Proxy     *httputil.ReverseProxy
	Addr      string
	IsHealthy bool

	// This will be a EMA
	AvgResponse int32
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func CreateReverseProxies(servers []Server) ([]Proxy, error) {
	var proxies []Proxy

	for i := range servers {
		url, err := url.Parse(servers[i].Addr)
		if err != nil {
			return nil, err
		}
		proxy := Proxy{
			Proxy:       NewProxy(url),
			Addr:        servers[i].Addr,
			IsHealthy:   true,
			AvgResponse: 1,
		}
		proxies = append(proxies, proxy)
	}
	return proxies, nil
}
