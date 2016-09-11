package main

import (
	"strconv"
	"time"
)

// ApaDest Holds the Craigs List URL
type ApaDest struct {
	City string
}

// FilterParams are parameters to filter results by
type FilterParams struct {
	RecentDate time.Time
	Price      int
	Location   []float64
}

// GetFilteredListings filters listings with filter params
// TODO: Replace filters with FetchParams
func GetFilteredListings(urlDest ApaDest, minDate time.Time, maxPrice int) (string, error) {
	url := urlDest.parseURL()
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
	return listing.Price < price
}

func getDesc(listing Listing) string {
	return ("Descr: " + listing.Desc +
		" | Price: " + strconv.Itoa(listing.Price) +
		"\nListing: " + listing.Link)
}

func (urlDest ApaDest) parseURL() string {
	return "http://" + urlDest.City + ".craiglist.org/search/apa"
}
