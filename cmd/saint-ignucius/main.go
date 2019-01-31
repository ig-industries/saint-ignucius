package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ig-industries/saint-ignucius/discord"
)

// Config represents a bot configuration. Ideally these values will be loaded
// in via environment variables, but since I don't trust Twitch, I'll keep this
// workflow.
type Config struct {
	Token string `json:"token,omitempty"`
	// Other config values...
}

func loadConfig() *Config {
	configFile, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("[main] %s", err)
	}

	config := &Config{}
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatalf("[main] %s", err)
	}

	return config
}

func main() {
	config := loadConfig()

	bot, err := discord.NewSession(config.Token)
	if err != nil {
		log.Fatalf("[main] %s", err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatalf("[main] %s", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	fmt.Println()
}
