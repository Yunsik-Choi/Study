package main

import (
	"OJT/network"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	ADDR = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.Path("/home").HandlerFunc(network.Home)
	router.Path("/login").HandlerFunc(network.LoginAuthentication).Methods(http.MethodGet)
	router.Path("/").HandlerFunc(network.IssueToken).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(ADDR, router))
}
