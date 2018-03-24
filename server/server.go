package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Sharykhin/gl-proxy/middleware"
	"github.com/Sharykhin/gl-proxy/proxy"
)

var (
	routes  map[string]string
	address string
)

const (
	routeFile = "routes.json"
)

func init() {
	raw, err := ioutil.ReadFile(routeFile)
	if err != nil {
		log.Fatalf("Could not read %s file: %v\n", routeFile, err)
	}
	err = json.Unmarshal(raw, &routes)
	if err != nil {
		log.Fatalf("Could not parse json %s file: %v\n", routeFile, err)
	}
	address = os.Getenv("HTTP_ADDRESS")
}

// Run starts listening and servers all income requests
func Run() error {
	p, err := proxy.NewProxy(routes)
	if err != nil {
		return err
	}
	server := &http.Server{
		Addr:    address,
		Handler: middleware.Chan(http.HandlerFunc(p.Handle), middleware.CORS),
	}

	fmt.Printf("Started listening on port %s\n", address)
	return server.ListenAndServe()
}
