// Code to scrape data from the Montreal Website

package main

import (
	"strings"

	"github.com/gocolly/colly"
)

// Pool struct
type pool struct {
	name          string
	href          string
	address       string
	neighbourhood string
}

// Stringer function for the pool struct
func (p pool) String() string {
	// Make a string with the pool name and href
	return "Name: " + p.name + "\nURL: " + p.href + "\nAddress: " + p.address + "\nNeighbourhood: " + p.neighbourhood
}

// Scraping the pool list with Colly
func getPoolList(url string) []pool {

	pools := make([]pool, 0)

	// Create a new instance of the Collector
	c := colly.NewCollector()

	c.OnHTML("a.list-group-item-action", func(e *colly.HTMLElement) {

		// Get the name of the pool with the text for only the specific div
		name := e.ChildText("div.list-group-item-title")

		// Get the URL of the pool
		href := e.Attr("href")

		// Get the address of the pool
		address := e.ChildText("div.list-group-item-infos.rm-last-child-mb")

		// Get the neighbourhood of the pool
		neighbourhood := strings.Replace(e.ChildText("div.list-group-item-infos"), address, "", 1)

		// Add the pool to the list
		pools = append(pools, pool{name, href, address, neighbourhood})
	})

	// Start scraping
	c.Visit(url)

	return pools

}
