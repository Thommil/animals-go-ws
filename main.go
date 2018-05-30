package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/thommil/animals-go-common/config"
	"github.com/thommil/animals-go-ws/middlewares"
	"github.com/thommil/animals-go-ws/resources/animals"
	"github.com/thommil/animals-go-ws/resources/users"
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

	Authentication middlewares.AuthenticationSettings
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	if err := config.LoadConfiguration("animals-go-ws", configuration); err != nil {
		log.Fatal(err)
	}

	//Mongo
	session, err := mgo.Dial(configuration.Mongo.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//HTTP Server
	router := gin.Default()

	//Middlewares
	router.Use(middlewares.Authenticated(&configuration.Authentication))

	//Resources
	users.New(router, session.DB(""))
	animals.New(router, session.DB(""))

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
