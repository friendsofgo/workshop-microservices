package main

import (
	"log"

	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/server/http"
)

const (
	_defaultHost      = "localhost"
	_defaultPort uint = 3000
)

func main() {
	srv := http.NewServer(_defaultHost, _defaultPort)
	log.Fatal(srv.Serve())
}
