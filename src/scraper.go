package main

import (
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

// GetListingsAfterDate returns entry descriptions that fall before a given date
func GetListingsAfterDate(url string, endDate time.Time) ([]Listing, error) {
	var recentListings []Listing
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
			recentListings = append(recentListings, currentListing)
		}
	})

	return recentListings, nil
}

/////////
// Listing helper functions:
/////////
func getImage(baseURL string, s *goquery.Selection) string {
	anchor := s.Find("a")
	imageLink, exists := anchor.Attr("href")
	if !exists {
		return "No Image Found."
	}
	return baseURL + imageLink
}

func createListing(s *goquery.Selection, date time.Time) Listing {
	return Listing{
		Date:  date,
		Desc:  s.Find(".txt .hdrlnk").Text(),
		Price: s.Find(".txt .price").Text(),
		Link:  getImage("http://slo.craiglist.org", s),
	}
}
