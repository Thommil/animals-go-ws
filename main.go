package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/config"
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
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	if err := config.LoadConfiguration("animals-go-ws", configuration); err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	router := gin.Default()

	//Resources
	users := users.New(router)
	animals := animals.New(router)

	fmt.Println(users, animals)

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
