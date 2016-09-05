package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ExtractHTML Reads row content from CL.
func ExtractHTML(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Get each CL entry
	doc.Find(".row").Each(formatText)
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

		// Parse CL datetime:
		format := "2006-01-02 15:04"
		date, _ := time.Parse(format, dt)
		today := time.Now()

		if InTimeSpan(endDate, today, date) {
			result += getDesc(s) + "\n\n"
		}
	})

	return result, nil
}

func formatText(i int, s *goquery.Selection) {
	// Posted Date
	postTime, _ := s.Find(".txt time").Attr("datetime")
	fmt.Println("\n" + postTime)

	// Extract description and price
	fmt.Println(getDesc(s))
}

func getImage(baseURL string, s *goquery.Selection) string {
	anchor := s.Find("a")
	imageLink, exists := anchor.Attr("href")
	if !exists {
		return "No Image Found."
	}
	return baseURL + imageLink
}

func getDesc(s *goquery.Selection) string {
	rowDesc := s.Find(".txt .hdrlnk").Text()
	rowPrice := s.Find(".txt .price").Text()
	imageAnchor := getImage("http://slo.craiglist.org", s)
	return "Descr: " + rowDesc + " | Price: " + rowPrice + "\nListing: " + imageAnchor
}
