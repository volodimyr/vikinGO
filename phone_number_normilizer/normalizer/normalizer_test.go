package normalizer

import (
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	tests := []struct {
		name string

		phone    string
		expected string
	}{
		{"Phone normalizer Test Case 1", "1234567890", "1234567890"},
		{"Phone normalizer Test Case 2", "123 456 7891", "1234567891"},
		{"Phone normalizer Test Case 3", "(123) 456 7892", "1234567892"},
		{"Phone normalizer Test Case 4", "(123) 456-7893", "1234567893"},
		{"Phone normalizer Test Case 5", "123-456-7894", "1234567894"},
		{"Phone normalizer Test Case 6", "123-456-7890", "1234567890"},
		{"Phone normalizer Test Case 7", "1234567892", "1234567892"},
		{"Phone normalizer Test Case 8", "(123)456-7892", "1234567892"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := phoneNumber(test.phone)
			if err != nil {
				t.Fatalf("Phone number func failed due to %v", err)
			}
			if actual == nil || len(actual) == 0 {
				t.Fatalf("Phone number func returned empty slice.")
			}
			if actual[0] != test.expected {
				t.Fatalf("Expected [%s] and actual [%s] don't match",
					test.expected, actual[0])
			}
		})
	}
}
