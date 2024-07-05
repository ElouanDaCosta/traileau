package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	auth, authErr := url.Parse("http://localhost:8081")
	if authErr != nil {
		panic(authErr)
	}
	authRedirectProxy := httputil.NewSingleHostReverseProxy(auth)

	project, projectErr := url.Parse("http://localhost:8082")
	if projectErr != nil {
		panic(projectErr)
	}
	projectRedirectProxy := httputil.NewSingleHostReverseProxy(project)

	http.Handle("/v1/auth/", authRedirectProxy)
	http.Handle("/v1/project/", projectRedirectProxy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
