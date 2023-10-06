package initializer

import (
	"log"

	"github.com/Zoopast/notes-app/utils"
)

func LoadConfig(configPath string) utils.Config {
	config, err := utils.LoadConfig(configPath)

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	return config
}
