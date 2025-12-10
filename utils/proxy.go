package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targerBaseURL string, pathPrefix string) http.HandlerFunc {

	target, err := url.Parse(targerBaseURL)

	if err != nil {
		fmt.Println("Error parsing url", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		fmt.Println("proxying req to: ", targerBaseURL)

		originalPath := r.URL.Path

		fmt.Println("original path: ", originalPath)

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)

		fmt.Println("stripped path: ", strippedPath)

		r.URL.Host = target.Host
		r.URL.Path = target.Path

		r.Host = target.Host
	}

	return proxy.ServeHTTP

	
}