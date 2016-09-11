package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(GetToken()) // Substitute with bot token
	params := slack.PostMessageParameters{}
	params.AsUser = true
	params.Username = "listingbot"

	filters := ReadSettingsFromFile("../settings.json")
	filters.RecentDate = time.Now().AddDate(0, 0, -filters.DaysPast)
	entries, err := GetFilteredListings(filters)
	if err != nil {
		log.Fatal("Could not fetch entries.")
	}

	channelID, timestamp, err := api.PostMessage("apartment", entries, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
