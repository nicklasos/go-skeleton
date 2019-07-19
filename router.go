package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routersServe() error {
	router := mux.NewRouter()

	if err := apiRouterInit(router); err != nil {
		return err
	}

	addrPort := os.Getenv("BIND_ADDRESS") + ":" + os.Getenv("PORT")

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})

	logger.Infoln("Starting server on", addrPort)

	if err := http.ListenAndServe(addrPort, handlers.CORS(origins, headers, methods)(router)); err != nil {
		return err
	}

	return nil
}
