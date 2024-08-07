package main

import (
	"fmt"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})

	return mux
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: Router(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
