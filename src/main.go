package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nlopes/slack"
)

func main() {
	config, err := ReadSettingsFromFile("settings.json")
	if err != nil {
		log.Fatal("Error reading settings file.", err.Error())
	}

	api := slack.New(config.Token) // Substitute with bot token
	params := slack.PostMessageParameters{}
	params.AsUser = true
	params.Username = "listingbot"

	config.RecentDate = time.Now().AddDate(0, 0, -config.DaysPast)
	entries, err := GetFilteredListings(config)
	if err != nil {
		log.Fatal("Could not fetch entries.")
	}

	channelID, timestamp, err := api.PostMessage("apartment", entries, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
