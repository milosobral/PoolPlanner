// Code to scrape data from the Montreal Website

package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Pool struct
type Pool struct {
	name          string
	href          string
	address       string
	neighbourhood string
}

type Schedule struct {
	validity []time.Time
	sched    map[string][]int
}

// define the priority of strings
var priorityList = []string{
	"la nage en couloir (16 ans et plus)",
	"la nage en couloir",
	"les adultes (16 ans et plus)",
	"les adultes",
	"toutes et tous (16 ans et plus)",
	"toutes et tous",
}

// Stringer function for the schedule struct
func (s Schedule) String() string {
	return fmt.Sprintf("Validity: %v\nSchedules: %v", s.validity, s.sched)
}

// Stringer function for the pool struct
func (p Pool) String() string {
	// Make a string with the pool name and href
	return "Name: " + p.name + "\nURL: " + p.href + "\nAddress: " + p.address + "\nNeighbourhood: " + p.neighbourhood
}

// Scraping the pool list with Colly
func getPoolList(url string) []Pool {

	pools := make([]Pool, 0)

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
		pools = append(pools, Pool{name, href, address, neighbourhood})
	})

	c.OnError(
		func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		},
	)

	// Start scraping
	c.Visit(url)

	return pools

}

// Takes a url and a channel and scrapes the pool schedule
func getPoolSchedule(url string) {

	schedules := make([]Schedule, 0)

	// Create a new instance of the Collector
	c := colly.NewCollector()

	c.OnHTML("div.wrapper[aria-label]", func(e *colly.HTMLElement) {
		selvalidity := e.DOM.Find("time[datetime]")

		// Get the validity of the schedule
		nodes := selvalidity.Nodes
		layout := "2006-01-02T15:04:05-0700"
		from, err := time.Parse(layout, nodes[0].Attr[0].Val)

		if err != nil {
			log.Fatal(err)
		}

		to, err := time.Parse(layout, nodes[0].Attr[0].Val)

		if err != nil {
			log.Fatal(err)
		}

		// Get the schedules for the current validity
		selsched := e.DOM.Find("div.content-module-stacked")

		// Parse the Schedule for the current validity
		poolSched := parseSchedule(selsched.Text())

		newSchedule := Schedule{
			validity: []time.Time{from, to},
			sched:    poolSched}
		schedules = append(schedules, newSchedule)

		fmt.Println(newSchedule)

	})

	// Start scraping
	c.Visit(url)
}

// Function to parse a schedule string and return a map of the days of the week
// with their corresponding hours of operation
func parseSchedule(scheduleString string) map[string][]int {

	// Split the schedule string into substrings that represent the different
	// days of the week and their corresponding hours of operation
	re := regexp.MustCompile("JourHoraire|Pour ")

	substrings := re.Split(scheduleString, -1)[1:]

	// Initialize a slice to store the priorities of the substrings
	priorities := make([]int, len(substrings)/2)

	// Loop over the substrings and set the priority of each substring
	// based on its index in the priority list
	for i, line := range substrings {
		if i%2 == 0 {
			// Set the priority based on the index of line in the priority list
			priorities[int(i/2)] = GetPriority(line)
		}
	}

	// Find the minimum priority
	minPriorityIndex := GetMinimumPriorityIndex(priorities)

	// Get the schedule for the minimum priority
	sched, err := GetScheduleFromString(substrings[minPriorityIndex*2+1])

	if err != nil {
		log.Fatal(err)
	}

	return sched

}

// Function to parse a schedule string and return a map of the days of the week
// with their corresponding hours of operation
func GetScheduleFromString(scheduleString string) (map[string][]int, error) {
	// Match the days of the week in french
	re_days := regexp.MustCompile("(Lundi|Mardi|Mercredi|Jeudi|Vendredi|Samedi|Dimanche)")

	// Split the string into days
	matches := re_days.FindAllStringSubmatchIndex(scheduleString, -1)

	result := make(map[string][]int)

	// Iterate over the matches and extract the content
	for i := 0; i < len(matches); i++ {
		day := scheduleString[matches[i][0]:matches[i][1]]

		var content string

		// Get the content of the day
		if i == len(matches)-1 {
			content = strings.TrimSpace(scheduleString[matches[i][1]:])
		} else {
			content = strings.TrimSpace(scheduleString[matches[i][1]:matches[i+1][0]])
		}

		// Remove all the white spaces from the content
		content = regexp.MustCompile(`\s|[\xa0]`).ReplaceAllString(content, "")

		// Add the content to the result
		re_hours := regexp.MustCompile(`h`)
		new_matches := re_hours.FindAllStringSubmatchIndex(content, -1)
		hours := make([]int, len(new_matches)*2)

		for i, match := range new_matches {
			var before, after string

			// Check the character before the match to know how many characters
			// to take for the hour
			if match[0] < 2 {
				before = content[match[0]-1 : match[0]]
			} else if match[0] == 2 {
				before = content[match[0]-2 : match[0]]
			} else {
				// Check if the hour is of the form 12h30 or 1h30
				test := content[match[0]-4 : match[0]]
				if regexp.MustCompile(`[^0-9].*\d{3}`).MatchString(test) {
					before = content[match[0]-1 : match[0]]
				} else {
					before = content[match[0]-2 : match[0]]
				}
			}

			if match[1] >= len(content) {
				return nil, errors.New("INVALID CONTENT")
			}

			after = content[match[1] : match[1]+2]

			// Convert string to int and store in hours
			beforeInt, err := strconv.Atoi(before)
			if err != nil {
				return nil, err
			}

			afterInt, err := strconv.Atoi(after)
			if err != nil {
				return nil, err
			}
			hours[i*2] = beforeInt
			hours[i*2+1] = afterInt

		}
		result[day] = hours

	}

	return result, nil
}

// Function to get the priority of a given substring
// The function takes a substring and returns its priority in the priority list
// If the substring is not found in the priority list, the function returns 1000
func GetPriority(substring string) int {
	// Iterate over the priority list
	for i, priority := range priorityList {
		// If the current priority matches the given substring, return its index
		if priority == substring {
			return i
		}
	}
	// If the substring is not found in the priority list, return 1000
	return 1000
}

// Function to find the index of the minimum priority in a list
// The function takes a list of priorities and returns the index of the minimum
// priority. If two priorities have the same value, the function returns the
// index of the first one.
func GetMinimumPriorityIndex(priorities []int) int {

	// Initialize the minimum priority and its index
	min := 1000
	var index int

	// Iterate over the list of priorities
	for i, priority := range priorities {
		// If the current priority is smaller than the minimum
		// update the minimum and its index
		if priority < min {
			min = priority
			index = i
		}
	}
	return index
}
