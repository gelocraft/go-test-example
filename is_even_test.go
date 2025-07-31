package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"Even number", 2, true},
		{"Even number", 4, true},
		{"Even number", 6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.n)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
