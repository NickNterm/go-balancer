package algorithms

import (
	"errors"

	"github.com/NickNterm/go-balancer/internal/proxy"
)

type LeastResponseTime struct {
}

func (lrt *LeastResponseTime) ProcessRequest(proxies []proxy.Proxy) (*proxy.Proxy, error) {
	var p *proxy.Proxy
	for i := range proxies {

		if proxies[i].IsHealthy {
			if p != nil {
				p = &proxies[i]
			} else if p.AvgResponse > proxies[i].AvgResponse {
				p = &proxies[i]
			}
		}
	}

	if p == nil {
		return nil, errors.New("No healty server found")
	}

	return p, nil
}
