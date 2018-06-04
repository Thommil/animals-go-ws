package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/config"
	"github.com/thommil/animals-go-common/dao/mongo"
	"github.com/thommil/animals-go-ws/middlewares/authentication"
	"github.com/thommil/animals-go-ws/resources/animals"
	"github.com/thommil/animals-go-ws/resources/users"
)

// Configuration definition for animals-go-ws
type Configuration struct {
	HTTP struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}

	Mongo mongo.Configuration

	Authentication authentication.Configuration
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	if err := config.LoadConfiguration("animals-go-ws", configuration); err != nil {
		log.Fatal(err)
	}

	//Mongo
	session, err := mongo.NewInstance(&configuration.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//HTTP Server
	router := gin.Default()

	//Middlewares
	router.Use(authentication.Authenticated(&configuration.Authentication))

	//Resources
	users.New(router)
	animals.New(router)

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
