package main

import (
	"fmt"
	//"github.com/thommil/animals-go-common/model"
	"github.com/thommil/animals-go-ws/config"
)

// Main of animals-go-ws
func main() {
	config, err := config.LoadConfiguration()

	if err != nil {
		panic(err)
	} else {
		fmt.Println(config.Mongo.Port)
	}

}
