package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nlopes/slack"
)

var (
	apartmentsURL = "http://slo.craiglist.org/search/apa"
)

func main() {
	api := slack.New(GetToken()) // Substitute with bot token
	params := slack.PostMessageParameters{}
	params.AsUser = true
	params.Username = "listingbot"

	oneDayAgo := time.Now().AddDate(0, 0, -1)
	entries, err := GetEntriesAfterDate(apartmentsURL, oneDayAgo)
	if err != nil {
		log.Fatal("Could not fetch entries.")
	}

	channelID, timestamp, err := api.PostMessage("apartment", entries, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
