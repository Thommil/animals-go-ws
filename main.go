package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thommil/animals-go-ws/api"
	"github.com/thommil/animals-go-ws/config"
	"log"
	"net/http"
	"strings"
	"time"
)

//All Resources (in package api) must implement this interface to allow subrouting
type Resource interface {
	GetRoutes(router *mux.Router) error
}

// Main of animals-go-ws
func main() {
	//Config
	config, err := config.LoadConfiguration()

	if err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", config.Http.Host, config.Http.Port)

	router := mux.NewRouter()
	server := &http.Server{
		Addr:         serverAddress.String(),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//Middlewares

	//users routes
	if err = users.GetRoutes(router.PathPrefix("/users").Subrouter()); err != nil {
		log.Fatal("Error applying users route", err)
	}

	//Start
	log.Printf("Starting HTTP server on %s\n", serverAddress.String())
	log.Fatal(server.ListenAndServe())
}
