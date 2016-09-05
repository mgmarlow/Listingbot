package main

import (
	"log"
	"time"

	"github.com/bluele/slack"
)

var (
	apartmentsURL = "http://slo.craiglist.org/search/apa"
)

func main() {
	api := slack.New(GetToken())

	oneDayAgo := time.Now().AddDate(0, 0, -1)
	entries, err := GetEntriesAfterDate(apartmentsURL, oneDayAgo)
	if err != nil {
		log.Fatal("Could not fetch entries.")
	}

	err = api.ChatPostMessage(
		"apartment",
		entries,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
