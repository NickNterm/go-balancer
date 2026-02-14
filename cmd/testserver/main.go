package main

import (
	"flag"
	"log"

	"github.com/NickNterm/go-balancer/internal/testserver"
)

func main() {

	port := flag.String("port", "8080", "server port")
	index := flag.String("index", "Hello world", "index text")

	flag.Parse()

	err := testserver.Start(*port, *index)
	if err != nil {
		log.Fatal(err)
	}

}
