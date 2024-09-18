// Test file to test the scraping code

package main

import "testing"

// Test for the getPoolList function
func TestGetPoolList(t *testing.T) {

	// Get the list of pools
	pools := getPoolList("https://montreal.ca/lieux?mtl_content.lieux.installation.code=PISI&mtl_content.lieux.available_activities.code=ACT0")

	// Check if the list contains at least one pool
	if len(pools) == 0 {
		t.Errorf("Expected the list to contain at least one pool, but it was empty")
	}

	// Check if the list contains at least one pool named Piscine Schubert
	var found bool = false
	for _, p := range pools {
		if p.name == "Piscine Schubert" {
			found = true
		}
	}

	if !found {
		t.Errorf("Expected the list to contain the pool named Piscine Schubert, but it was not found")
	}

}
