// Test file to test the scraping code

package scraping

import (
	"fmt"
	"reflect"
	"testing"
)

// Test for the getPoolList function
func TestGetPoolList(t *testing.T) {

	// Get the list of pools
	pools := GetPoolList("https://montreal.ca/lieux?mtl_content.lieux.installation.code=PISI&mtl_content.lieux.available_activities.code=ACT0")

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

func TestGetScheduleFromString(t *testing.T) {
	// Test case 1
	scheduleString := "Lundi8h30a12h3013h30a16h30"
	expected := map[string][]int{"Lundi": {8, 30, 12, 30, 13, 30, 16, 30}}
	actual, err := GetScheduleFromString(scheduleString)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test case 2
	scheduleString = "Lundi8h30a12h303h30a16h30"
	expected = map[string][]int{"Lundi": {8, 30, 12, 30, 3, 30, 16, 30}}
	actual, err = GetScheduleFromString(scheduleString)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test case 3
	scheduleString = "Lundi8h30a12h303h30a16h30Mardi9h35a13h35"
	expected = map[string][]int{"Lundi": {8, 30, 12, 30, 3, 30, 16, 30}, "Mardi": {9, 35, 13, 35}}
	actual, err = GetScheduleFromString(scheduleString)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test Case 4
	scheduleString = "Lundi8h30a12h303h30a16h30Mardi9h35a13h35Samedi8h30a11h30"
	expected = map[string][]int{"Lundi": {8, 30, 12, 30, 3, 30, 16, 30}, "Mardi": {9, 35, 13, 35}, "Samedi": {8, 30, 11, 30}}
	actual, err = GetScheduleFromString(scheduleString)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test case 5
	scheduleString = ""
	expected = map[string][]int{}
	actual, err = GetScheduleFromString(scheduleString)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test case 6
	scheduleString = "Samedi8h"
	expected = nil
	actual, err = GetScheduleFromString(scheduleString)
	fmt.Println(actual)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}

// Test case for the GetPriority function
func TestGetPriority(t *testing.T) {
	// Test case 1
	substring := "la nage en couloir (16 ans et plus)"
	expected := 0
	actual := GetPriority(substring)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	// Test case 2
	substring = "les adultes (16 ans et plus)"
	expected = 2
	actual = GetPriority(substring)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	// Test case 3
	substring = "foo"
	expected = 1000
	actual = GetPriority(substring)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

// Test for the GetMinimumPriorityIndex function
func TestGetMinimumPriorityIndex(t *testing.T) {
	// Test case 1
	priorities := []int{1, 2, 3}
	expected := 0
	actual := GetMinimumPriorityIndex(priorities)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	// Test case 2
	priorities = []int{10, 20, 30}
	expected = 0
	actual = GetMinimumPriorityIndex(priorities)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	// Test case 3
	priorities = []int{30, 20, 10}
	expected = 2
	actual = GetMinimumPriorityIndex(priorities)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
