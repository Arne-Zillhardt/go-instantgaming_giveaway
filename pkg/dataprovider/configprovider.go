package dataprovider

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type data struct {
	Alive []string `json:"alive"`
	Dead  []string `json:"dead"`
	BrowserXPosition string `json:"x-position-browser"`
	BrowserYPosition string `json:"y-position-browser"`
	ParticipationXPosition string `json:"x-position-participate"`
	ParticipationYPosition string `json:"y-position-participate"`
	ParticipationCheckXPosition string `json:"x-position-participation-check"`
	ParticipationCheckYPosition string `json:"y-position-participation-check"`
	ExtrasXPosition string `json:"x-position-extras"`
	ExtrasYStartPosition string `json:"y-position-extras-start"`
	ExtrasYEndPosition string `json:"y-position-extras-end"`
}

var configuration data

func getConfigData() data {
	if configuration.ExtrasXPosition != "" {
		return configuration
	}

	file, err := os.Open("./assets/config.json")
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var payload data
	err = json.NewDecoder(reader).Decode(&payload)
	if err != nil {
		log.Fatal("Error while parsing file: ", err)
	}

	configuration = payload
	return payload
}

func GetBrowserPositionX() string {
	return configuration.BrowserXPosition
}

func GetBrowserPositionY() string {
	return configuration.BrowserYPosition
}

func GetParticipationXPosition() string {
	return configuration.ParticipationXPosition
}

func GetParticipationYPosition() string {
	return configuration.ParticipationYPosition
}

func GetParticipationCheckXPosition() string {
	return configuration.ParticipationCheckXPosition
}

func GetParticipationCheckYPosition() string {
	return configuration.ParticipationCheckYPosition
}

func GetExtrasXPosition() string {
	return configuration.ExtrasXPosition
}

func GetExtrasYPositionStart() string {
	return configuration.ExtrasYStartPosition
}

func GetExtrasYPositionEnd() string {
	return configuration.ExtrasYEndPosition
}
