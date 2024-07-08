package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(rawUrl string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	return proxy, nil
}

func main() {
	auth, authErr := NewProxy("http://localhost:8081")
	if authErr != nil {
		panic(authErr)
	}

	project, projectErr := NewProxy("http://localhost:8082")
	if projectErr != nil {
		panic(projectErr)
	}

	http.Handle("/v1/auth/", auth)
	http.Handle("/v1/project/", project)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
