package testserver

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Start(port string, index string) error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration := rand.Intn(100) + 100
		time.Sleep(time.Duration(duration) * time.Millisecond)
		w.Write([]byte(index))
	})

	log.Println("Running on port", port)

	return http.ListenAndServe(":"+port, nil)
}
