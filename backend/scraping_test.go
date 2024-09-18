// Test file to test the scraping code

package main

import (
	"testing"
)

// Test case for getURL
func TestGetHTML(t *testing.T) {
	url := "https://montreal.ca/lieux/piscine-schubert"
	htmlCode, err := getHTML(url)
	if err != nil {
		t.Errorf("getHTML(%s) returned an error: %s", url, err)
	}
	if len(htmlCode) == 0 {
		t.Errorf("getHTML(%s) returned no HTML code", url)
	}
}
