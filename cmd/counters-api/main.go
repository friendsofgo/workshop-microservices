package main

import (
	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/server/http"
)

const (
	_defaultHost       = "localhost"
	_defaultPort uint = 3000
)

func main() {
	srv := http.NewServer(_defaultHost, _defaultPort)
	srv.Serve()
}
