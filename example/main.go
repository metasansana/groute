package main

import (
	"fmt"
	"github.com/metasansana/groute"
	"net/http"
)

func main() {

	router := groute.MakeRouter()

	router.Get(router.Static("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("/ route accessed!")

	})).
		Get(router.Static("/secret", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("ssssssssh!! It's a secret don't tell anyone!")

	}))

	http.ListenAndServe(":8080", router)

}
