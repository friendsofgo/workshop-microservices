package main

import (
	"log"
	"os"

	"github.com/friendsofgo/workshop-microservices/cmd/counters-api/server/http"
)

const (
	_defaultHost      = "localhost"
	_defaultPort uint = 3000
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	srv := http.NewServer(_defaultHost, _defaultPort, logger)
	log.Fatal(srv.Serve())
}
