package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func loadConfigFromFile(filePath string) (*configData, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening the file: error = %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config configData

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Unable to unmarshall file: error = %v", err)
	}
	return &config, nil
}
