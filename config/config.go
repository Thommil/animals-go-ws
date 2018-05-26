package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path/filepath"
	"strings"
)

// Configuration definition for animals-go-ws
type Configuration struct {
	Mongo struct {
		Host string
		Port int
	}
}

// Load configuration from json file to Configuration
func LoadConfiguration() (Configuration, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	var filePath strings.Builder
	fmt.Fprintf(&filePath, "%s/animals-go-ws.json", exPath)
	configuration := Configuration{}
	err = gonfig.GetConf(filePath.String(), &configuration)
	return configuration, err
}
