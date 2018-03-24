package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
)

type (
	// Proxy is a struct that handles income request and proxies them.
	// It keep routes to know which request should be sent to.
	Proxy struct {
		routes map[*regexp.Regexp]*httputil.ReverseProxy
	}
)

var (
	corsOrigin string
)

func init() {
	corsOrigin = os.Getenv("CORS_ORIGIN")
}

// NewProxy returns a new proxy with registered path and appropriate destinations
// currently it accepts the following map: path -> server
// Example:
// map[string]string{
//  "^/users*|/register|/login": "http://127.0.0.1:8080",
//  "^/maps*" : "http://127.0.0.1:8081",
// }
func NewProxy(servers map[string]string) (*Proxy, error) {
	routes := make(map[*regexp.Regexp]*httputil.ReverseProxy)

	for pattern, target := range servers {
		targetURL, err := url.ParseRequestURI(target)
		if err != nil {
			return nil, fmt.Errorf("could not parse a provided url %s: %v", target, err)
		}
		c, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("could not compile a regular expression %s: %v", pattern, err)
		}
		routes[c] = httputil.NewSingleHostReverseProxy(targetURL)
	}
	return &Proxy{routes: routes}, nil
}

func (p *Proxy) parseRequest(r *http.Request) *httputil.ReverseProxy {
	for regExpCompile, proxyServer := range p.routes {
		fmt.Println(r.URL.Path, regExpCompile.String())
		if regExpCompile.MatchString(r.URL.Path) {
			return proxyServer
		}
	}
	return nil
}

// Handle handles all income requests.
// It fills it up with some additional headers and pass to a target server
func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	// TODO: Move CORS to a separate middleware
	if r.Method == "OPTIONS" {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	// TODO: Move Additional headers to a separate middleware
	w.Header().Set("X-GoProxy", "GoProxy")
	w.Header().Set("X-Forwarded-Proto", "http")
	w.Header().Set("X-Real-IP", r.RemoteAddr)
	//w.Header().Set("X-Forwarded-For", p.target.String())

	if proxyServer := p.parseRequest(r); proxyServer != nil {
		proxyServer.ServeHTTP(w, r)
	} else {
		http.Error(w, "Route not found", http.StatusNotFound)
	}
}
