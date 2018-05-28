package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/config"
	"github.com/thommil/animals-go-ws/api"
)

// Configuration definition for animals-go-ws
type Configuration struct {
	HTTP struct {
		Host string
		Port int
	}

	Mongo struct {
		URL string
	}
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	err := config.LoadConfiguration("animals-go-ws", configuration)

	if err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	router := gin.Default()

	//users API
	users := &api.Users{Engine: router}
	users.ApplyRoutes()

	//animals API
	animals := &api.Animals{Engine: router}
	animals.ApplyRoutes()

	//Middlewares

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Printf("Starting HTTP server on %s\n", serverAddress.String())
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
