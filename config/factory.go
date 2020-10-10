package config

import (
	"log"

	"github.com/gorilla/mux"
)

var data *configData

const filePath = "config/config.json"

// HTTPConfig interface
type HTTPConfig interface {
	GetPort() int
}

// Repository interface
type Repository interface {
	GetPort() int
	GetPublicKey() string
	GetPrivateKey() string
	GetSigningAlgo() string
}

func init() {
	var err error
	data, err = loadConfigFromFile(filePath)
	if err != nil {
		log.Fatalf("Unable to load config from %s: error = %v", filePath, err)
	}
}

// GetHTTPConfig returns HTTPConfig interface
func GetHTTPConfig() HTTPConfig {
	return data
}

// GetRepo returns repo interface
func GetRepo() Repository {
	return data
}

// InitService initialises the Config Service
func InitService(router *mux.Router) {
	cs := newConfigService(data)
	cs.AttachRoute(router)
}
