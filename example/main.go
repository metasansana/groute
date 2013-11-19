package main

import (
	"fmt"
	"github.com/metasansana/groute"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[%s] request detected on path %s.", r.Method, r.URL.Path)

}

func main() {

	router := groute.MakeRouter()

	router.Get(router.Static("/static/get", handler)).
		Post(router.Static("/static/post", handler)).
		Put(router.Static("/static/put", handler)).
		Head(router.Static("/static/head", handler)).
		Delete(router.Static("/static/delete", handler)).
		Get(router.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Post(router.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Put(router.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Head(router.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Delete(router.Regex("/regex/[a-zA-Z0-9]+", handler))

	http.ListenAndServe(":8080", router)

}
