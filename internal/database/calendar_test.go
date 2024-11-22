// Test files for the database connection for the Calendar Table
package database_test

import (
	"testing"

	"github.com/milosobral/PoolPlanner/internal/database"
)

func TestGenerateUniqueHash(t *testing.T) {
	// Test the generateUniqueHash function
	tests := []struct {
		name  string
		input []string
		wants string
	}{
		{
			name:  "Testing with valid input",
			input: []string{"Name1", "Name2"},
			wants: "eba722137a",
		},
		{
			name:  "Testing with a single input",
			input: []string{"Name2"},
			wants: "f68b8f97da",
		},
		{
			name:  "Testing with an empty string",
			input: []string{""},
			wants: "e3b0c44298",
		},
		{
			name:  "Testing with an empty slice",
			input: []string{},
			wants: "e3b0c44298",
		},
	}

	// Run the tests
	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				got := db.GenerateUniqueHash(test.input)
				if got != test.wants {
					t.Errorf("generateUniqueHash() = %v, want %v", got, test.wants)
				}
			},
		)
	}

}
