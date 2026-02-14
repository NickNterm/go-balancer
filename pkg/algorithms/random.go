package algorithms

import (
	"errors"
	"math/rand"

	"github.com/NickNterm/go-balancer/internal/proxy"
)

type Random struct {
}

func (r *Random) ProcessRequest(proxies []proxy.Proxy) (*proxy.Proxy, error) {
	var index int
	counter := 0
	for true {
		counter++
		index = rand.Intn(len(proxies))
		if proxies[index].IsHealthy {
			return &proxies[index], nil
		}
		if counter > 1000 {
			return nil, errors.New("No healthy server found")
		}
	}
	return nil, nil
}
