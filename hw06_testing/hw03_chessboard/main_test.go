package main

import (
	"testing"
)

func TestGetChessboerd(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected string
	}{
		{"1", 1, "#\r\n"},
		{"3", 3, "# #\r\n # \r\n# #\r\n"},
		{"4", 4, "# # \r\n # #\r\n# # \r\n # #\r\n"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := GetChessboerd(tc.size)
			if got != tc.expected {
				t.Errorf("size(%q) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}
