// Code to scrape data from the Montreal Website

package main

import (
	"fmt"
	"regexp"
	"strings"

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
	validity string
	sched    map[string]string
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

	// Create a new instance of the Collector
	c := colly.NewCollector()

	c.OnHTML("div.wrapper[aria-label]", func(e *colly.HTMLElement) {
		selvalidity := e.DOM.Find("time[datetime]")

		// Get the validity of the schedule
		nodes := selvalidity.Nodes
		for _, node := range nodes {
			fmt.Println(node.Attr)
		}

		// Get the schedules for the current validity
		selsched := e.DOM.Find("div.content-module-stacked")
		// fmt.Println(selsched.Text())

		// Parse the Schedule for the current validity
		parseSchedule(selsched.Text())

	})

	// Start scraping
	c.Visit(url)
}

func parseSchedule(scheduleString string) map[string][]string {

	re := regexp.MustCompile("JourHoraire|Pour ")

	substrings := re.Split(scheduleString, -1)[1:]

	priorities := make([]int, len(substrings)/2)

	for i, line := range substrings {
		if i%2 == 0 {
			// Set the priority based on the index of line in the priority list
			priorities[int(i/2)] = GetPriority(line)
		}
	}

	// Find the minimum priority
	minPriorityIndex := GetMinimumPriorityIndex(priorities)

	// Get the schedule for the minimum priority
	sched := GetScheduleFromString(substrings[minPriorityIndex*2+1])

	return sched

}

func GetScheduleFromString(scheduleString string) map[string][]string {
	// Match the days of the week in french
	re_days := regexp.MustCompile("(Lundi|Mardi|Mercredi|Jeudi|Vendredi|Samedi|Dimanche)")

	// Split the string into days
	matches := re_days.FindAllStringSubmatchIndex(scheduleString, -1)

	result := make(map[string]string)

	// Iterate over the matches and extract the content
	for i := 0; i < len(matches)-1; i++ {
		day := scheduleString[matches[i][0]:matches[i][1]]
		content := strings.TrimSpace(scheduleString[matches[i][1]:matches[i+1][0]])

		// Remove all the white spaces from the content
		content = strings.ReplaceAll(content, " ", "")

		// Add the content to the result
		pattern := regexp.MustCompile(`\d{0,2}h\d{0,2}`)
		new_matches := pattern.FindAllString(content, -1)

		for _, new_match := range new_matches {
			fmt.Println(new_match)
		}

		result[day] = content

	}

	return nil
}

func GetPriority(substring string) int {
	for i, priority := range priorityList {
		if priority == substring {
			return i
		}
	}
	return 1000
}

func GetMinimumPriorityIndex(priorities []int) int {

	min := 1000
	var index int

	for i, priority := range priorities {
		if priority < min {
			min = priority
			index = i
		}
	}
	return index
}

// func extractTimes(input string) []map[string]string {

// 	//

// 	return info
// }

// func parseSchedule(scheduleString string) {
// 	// Regular expressions to match different parts of the schedule
// 	periodRegex := regexp.MustCompile(`du (\d{2} au \d{2})`)
// 	// dayRegex := regexp.MustCompile(`(Lundi|Mardi|Mercredi|Jeudi|Vendredi|Samedi|Dimanche)`)
// 	// timeRegex := regexp.MustCompile(`(\d{2} h \d{2} à \d{2} h \d{2})`)

// 	// schedule := make(Schedule)
// 	// currentPeriod := ""

// 	for _, line := range strings.Split(scheduleString, "\n") {
// 		fmt.Println(line)
// 		fmt.Println()
// 		// Match the period
// 		// if periodMatch := periodRegex.FindStringSubmatch(line); len(periodMatch) > 1 {
// 		// 	currentPeriod = periodMatch[1]
// 		// 	// schedule[currentPeriod] = make(map[string][]string)

// 		// 	fmt.Println(currentPeriod)
// 		// }

// 		// Match the day and time
// 		// if dayMatch := dayRegex.FindStringSubmatch(line); len(dayMatch) > 0 {
// 		// 	day := dayMatch[1]
// 		// 	timeMatches := timeRegex.FindAllStringSubmatch(line, -1)
// 		// 	times := make([]string, len(timeMatches))
// 		// 	for i, match := range timeMatches {
// 		// 		times[i] = match[1]
// 		// 	}
// 		// 	// schedule[currentPeriod][day] = times
// 		// }
// 	}

// 	// return schedule
// }
