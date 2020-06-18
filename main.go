package main

import (
	"log"

	"github.com/ariefsam/go-rest-api-proxy/delivery/rest"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rest.Serve()
}
