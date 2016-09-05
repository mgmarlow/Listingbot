package main

import (
	"fmt"
	"log"

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

func formatText(i int, s *goquery.Selection) {
	// Posted Date
	postTime, _ := s.Find(".txt time").Attr("datetime")
	fmt.Println("\n" + postTime)

	// Extract description and price
	rowDesc := s.Find(".txt .hdrlnk").Text()
	rowPrice := s.Find(".txt .price").Text()
	fmt.Println("Descr:", rowDesc, "Price:", rowPrice)

	// CL Link
	imageAnchor := getImage("http://slo.craiglist.org", s)
	fmt.Println("Listing:", imageAnchor)
}

func getImage(baseURL string, s *goquery.Selection) string {
	imageAnchor, exists := s.Find("a").Attr("href")
	i, _ := s.Find("a").Attr("title")
	if !exists || i == "no image" {
		return "No Image Found."
	}
	return baseURL + imageAnchor
}
