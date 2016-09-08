package main

import (
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Listing holds apartment listing data
type Listing struct {
	Desc  string
	Price string
	Link  string
}

// GetEntriesAfterDate returns entry descriptions that fall before a given date
func GetEntriesAfterDate(url string, endDate time.Time) (string, error) {
	var result string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		dt, _ := s.Find(".txt time").Attr("datetime")
		currentListing := createListing(s)

		// Parse CL datetime:
		format := "2006-01-02 15:04"
		date, _ := time.Parse(format, dt)
		today := time.Now()

		if InTimeSpan(endDate, today, date) {
			result += getDesc(currentListing) + "\n\n"
		}
	})

	return result, nil
}

func getImage(baseURL string, s *goquery.Selection) string {
	anchor := s.Find("a")
	imageLink, exists := anchor.Attr("href")
	if !exists {
		return "No Image Found."
	}
	return baseURL + imageLink
}

func createListing(s *goquery.Selection) Listing {
	newListing := Listing{
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
