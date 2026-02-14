package healtchecker

import (
	"net/http"
	"time"
)

func CheckHealth(url string) (bool, int32) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return false, 1
	}
	end := time.Now()

	return true, int32(end.Sub(start).Milliseconds())
}
