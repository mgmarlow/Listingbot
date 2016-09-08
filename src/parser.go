package main

import (
	"log"
	"strconv"
	"time"
)

// FilterParams are parameters to filter results by
type FilterParams struct {
	RecentDate time.Time
	Price      int
	Location   []float64
}

// GetFilteredListings filters listings with filter params
// TODO: Replace filters with FetchParams
func GetFilteredListings(url string, minDate time.Time, maxPrice int) (string, error) {
	var result string
	listings, err := GetListingsAfterDate(url, minDate)
	if err != nil {
		return "", err
	}

	for _, listing := range listings {
		if listing.withinBudget(maxPrice) {
			result += getDesc(listing) + "\n\n"
		}
	}

	return result, nil
}

func (listing Listing) withinBudget(price int) bool {
	listingPrice, err := strconv.Atoi(listing.Price[1:])
	if err != nil {
		log.Fatal("ERror converting price to integer.")
	}
	return listingPrice < price
}

func getDesc(listing Listing) string {
	return ("Descr: " + listing.Desc +
		" | Price: " + listing.Price +
		"\nListing: " + listing.Link)
}
