package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targerBaseURL string, pathPrefix string) http.HandlerFunc {
					// ("https://fakestoreapi.com/"  ,  "/fakestore")
	target, err := url.Parse(targerBaseURL)	// a url object is created

	/**
		{
			Scheme: "https",
			Host: "fakestoreapi.com",
			Path: "",
			RawQuery: ""
		}
	*/

	if err != nil {
		fmt.Println("Error parsing url", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)	// proxy created

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {// director defines how the req is going to be modified before its sent to the target backend
		originalDirector(r)
		// request came as (r) --> http://localhost:3001/fakestore/users/1

		fmt.Println("proxying req to: ", targerBaseURL)
		//	"https://fakestoreapi.com/"	<-- tagetBaseURL

		originalPath := r.URL.Path
		// /fakestore/users/1 <-- originalPath

		fmt.Println("original path: ", originalPath)

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)
		// /users/1  <-- strippedPath

		fmt.Println("stripped path: ", strippedPath)

		r.URL.Host = target.Host	
		/*
			{
				Host: "fakestoreapi.com"
			}
		*/
		r.URL.Path = strippedPath
		/*
			{
				Host: "fakestoreapi.com",
				Path: "/users/1"
			}
		*/

		r.Host = target.Host
	}

	return proxy.ServeHTTP

	
}