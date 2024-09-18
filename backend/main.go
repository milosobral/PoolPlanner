package main

import (
	"fmt"
	"log"
)

func main() {
	html_from_page, err := getHTML("https://montreal.ca/lieux/piscine-schubert")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html_from_page)
}
