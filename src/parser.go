package main

import (
	"log"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Listing holds apartment listing data
type Listing struct {
	Date  time.Time
	Desc  string
	Price string
	Link  string
}

// FetchParams are parameters to filter results by
type FetchParams struct {
	RecentDate time.Time
	Price      int
	Location   []float64
}

// GetAdequateListings filters listings with filter params
// TODO: Replace filters with FetchParams
func GetAdequateListings(url string, minDate time.Time, maxPrice int) (string, error) {
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

// GetListingsAfterDate returns entry descriptions that fall before a given date
func GetListingsAfterDate(url string, endDate time.Time) ([]Listing, error) {
	var result []Listing
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		dt, _ := s.Find(".txt time").Attr("datetime")
		// Parse CL datetime:
		format := "2006-01-02 15:04"
		date, _ := time.Parse(format, dt)

		currentListing := createListing(s, date)
		today := time.Now()

		if InTimeSpan(endDate, today, date) {
			result = append(result, currentListing)
		}
	})

	return result, nil
}

func (listing Listing) withinBudget(price int) bool {
	listingPrice, err := strconv.Atoi(listing.Price[1:])
	if err != nil {
		log.Fatal("ERror converting price to integer.")
	}
	return listingPrice < price
}

func getImage(baseURL string, s *goquery.Selection) string {
	anchor := s.Find("a")
	imageLink, exists := anchor.Attr("href")
	if !exists {
		return "No Image Found."
	}
	return baseURL + imageLink
}

// TODO: separate listing and scraping logic
func createListing(s *goquery.Selection, date time.Time) Listing {
	newListing := Listing{
		Date:  date,
		Desc:  s.Find(".txt .hdrlnk").Text(),
		Price: s.Find(".txt .price").Text(),
		Link:  getImage("http://slo.craiglist.org", s),
	}
	return newListing
}

func getDesc(listing Listing) string {
	return ("Descr: " + listing.Desc +
		" | Price: " + listing.Price +
		"\nListing: " + listing.Link)
}
