package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SimpleProxy struct {
	Proxy *httputil.ReverseProxy
}

func NewProxy(rawUrl string) (*SimpleProxy, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	s := &SimpleProxy{httputil.NewSingleHostReverseProxy(url)}

	// Modify requests
	// originalDirector := s.Proxy.Director
	// s.Proxy.Director = func(r *http.Request) {
	// 	originalDirector(r)
	// 	r.Header.Set("Some-Header", "Some Value")
	// }

	// // Modify response
	// s.Proxy.ModifyResponse = func(r *http.Response) error {
	// 	// Add a response header
	// 	r.Header.Set("Server", "CodeDodle")
	// 	return nil
	// }

	return s, nil
}

func (s *SimpleProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Do anything you want here
	// e.g. blacklisting IP, log time, modify headers, etc
	log.Printf("Proxy receives request from.")
	log.Printf("Proxy forwards request to ")
	s.Proxy.ServeHTTP(w, r)
	log.Printf("Origin server completes request.")
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
