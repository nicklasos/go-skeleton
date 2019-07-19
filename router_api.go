package main

import "github.com/gorilla/mux"

func apiRouterInit(router *mux.Router) error {
	router.HandleFunc("/test", testHandler).Methods("GET")

	return nil
}
