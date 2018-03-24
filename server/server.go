package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Sharykhin/gl-proxy/proxy"
)

var (
	routes  map[string]string
	address string
)

const (
	ROUTE_FILE = "routes.json"
)

func init() {
	raw, err := ioutil.ReadFile(ROUTE_FILE)
	if err != nil {
		log.Fatalf("Could not read %s file: %v\n", ROUTE_FILE, err)
	}
	json.Unmarshal(raw, &routes)
	address = os.Getenv("HTTP_ADDRESS")
}

// Run starts listening and servers all income requests
func Run() error {
	p := proxy.NewProxy(routes)
	server := &http.Server{
		Addr:    address,
		Handler: http.HandlerFunc(p.Handle),
	}

	fmt.Printf("Started listening on port %s\n", address)
	return server.ListenAndServe()
}
