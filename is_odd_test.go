package main

import "testing"

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"Odd number", 1, true},
		{"Odd number", 3, true},
		{"Odd number", 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOdd(tt.n)
			if result != tt.expected {
				t.Errorf("IsOdd(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
