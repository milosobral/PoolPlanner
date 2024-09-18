// Code to scrape data from the Montreal Website

package main

import (
	"io"
	"net/http"
)

// Function to make a request to the website and get the data
func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Read the HTML content as a byte slice
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	htmlCode := string(htmlBytes)

	return htmlCode, nil
}
