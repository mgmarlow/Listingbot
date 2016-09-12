package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

// FilterParams are parameters to filter results by
type FilterParams struct {
	Token      string    `json:"slackToken"`
	City       string    `json:"city"`
	RecentDate time.Time `json:"recentDate"`
	DaysPast   int       `json:"daysPast"`
	Price      int       `json:"price"`
	Location   []float64 `json:"location"`
}

// ReadSettingsFromFile fetches JSON settings data
func ReadSettingsFromFile(fileName string) (FilterParams, error) {
	var filters FilterParams
	configFile, err := os.Open(fileName)
	if err != nil {
		return FilterParams{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&filters)
	if err != nil {
		return FilterParams{}, err
	}

	return filters, nil
}

// GetFilteredListings filters listings with filter params
func GetFilteredListings(filters FilterParams) (string, error) {
	url := filters.parseURL()
	var result string
	listings, err := GetListingsAfterDate(url, filters.RecentDate)
	if err != nil {
		return "", err
	}

	for _, listing := range listings {
		if listing.withinBudget(filters.Price) {
			result += getDesc(listing) + "\n\n"
		}
	}

	return result, nil
}

func (listing Listing) withinBudget(price int) bool {
	return listing.Price < price
}

func getDesc(listing Listing) string {
	return ("Descr: " + listing.Desc +
		" | Price: " + strconv.Itoa(listing.Price) +
		"\nListing: " + listing.Link)
}

func (urlDest FilterParams) parseURL() string {
	return "http://" + urlDest.City + ".craiglist.org/search/apa"
}
