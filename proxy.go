package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type Proxy struct {
	route map[*regexp.Regexp]*httputil.ReverseProxy
}

// NewProxy returns a new proxy with registered path and appropriate destinations
// currently it accepts the following map: path -> server
// Example:
// map[string]string{
//  "users": "http://127.0.0.1:8080",
//  "maps" : "http://127.0.0.1:8081",
// }
func NewProxy(servers map[string]string) *Proxy {
	route := make(map[*regexp.Regexp]*httputil.ReverseProxy)

	for pattern, target := range servers {
		targetUrl, err := url.Parse(target)
		if err != nil {
			log.Fatalf("Could not parse a provided url %s: %v\n", target, err)
		}
		c, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatalf("Could not compile a regular expression %s: %v\n", pattern, err)
		}
		route[c] = httputil.NewSingleHostReverseProxy(targetUrl)
	}
	return &Proxy{route: route}
}

func (p *Proxy) parseRequest(r *http.Request) *httputil.ReverseProxy {
	for regExpCompile, proxyServer := range p.route {
		fmt.Println(r.URL.Path, regExpCompile.String())
		if regExpCompile.MatchString(r.URL.Path) {
			return proxyServer
		}
	}
	return nil
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-GoProxy", "GoProxy")
	w.Header().Set("X-Forwarded-Proto", "http")
	w.Header().Set("X-Real-IP", r.RemoteAddr)
	//w.Header().Set("X-Forwarded-For", p.target.String())

	if proxyServer := p.parseRequest(r); proxyServer != nil {
		proxyServer.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route not found"))
	}
}

var (
	routes map[string]string
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
}

func main() {
	proxy := NewProxy(routes)
	http.HandleFunc("/", proxy.handle)
	fmt.Println("Started listening on port :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
