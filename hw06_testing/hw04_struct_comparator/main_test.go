package main

import (
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		c        Comparator
		b1       Book
		b2       Book
		expected bool
	}{
		{"year 1", *NewComparator(1), *NewBook(1, "1", "1", 1995, 200, 5), *NewBook(2, "2", "2", 1996, 150, 10), false},
		{"year 2", *NewComparator(1), *NewBook(1, "1", "1", 1998, 200, 5), *NewBook(2, "2", "2", 1996, 150, 10), true},
		{"size 1", *NewComparator(2), *NewBook(1, "1", "1", 1998, 200, 5), *NewBook(2, "2", "2", 1996, 150, 10), true},
		{"size 2", *NewComparator(2), *NewBook(1, "1", "1", 1998, 200, 5), *NewBook(2, "2", "2", 1996, 250, 10), false},
		{"rete 1", *NewComparator(3), *NewBook(1, "1", "1", 1998, 200, 5), *NewBook(2, "2", "2", 1996, 150, 10), false},
		{"rete 2", *NewComparator(3), *NewBook(1, "1", "1", 1998, 200, 5), *NewBook(2, "2", "2", 1996, 150, 1), true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.Compare(&tc.b1, &tc.b2)
			if got != tc.expected {
				t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}
