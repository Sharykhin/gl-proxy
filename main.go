package main

import (
	"log"

	"github.com/Sharykhin/gl-proxy/server"
)

func main() {
	log.Fatal(server.Run())
}
