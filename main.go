package main

import (
	"fmt"
	"github.com/thommil/animals-go-common/model"
	"github.com/thommil/animals-go-ws/config"
	"log"
)

//Mongo : https://godoc.org/github.com/globalsign/mgo
//MUX : http://www.gorillatoolkit.org/pkg/mux

// Main of animals-go-ws
func main() {
	config, err := config.LoadConfiguration()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(config.Mongo.Url)
		fmt.Println(user.Get())
	}

}
