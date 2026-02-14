package algorithms

import (
	"errors"

	"github.com/NickNterm/go-balancer/internal/proxy"
)

type RoundRobin struct {
	CurrentIndex int
}

func (rrd *RoundRobin) ProcessRequest(proxies []proxy.Proxy) (*proxy.Proxy, error) {
	index := rrd.CurrentIndex
	var p *proxy.Proxy
	for range proxies {
		p = &proxies[index]
		index++
		index %= len(proxies)
		if p.IsHealthy {
			break
		}
	}
	if !p.IsHealthy {
		return nil, errors.New("No healthy server found")
	}
	rrd.CurrentIndex = index
	return p, nil
}
