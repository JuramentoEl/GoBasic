package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		s        []int
		i        int
		expected int
	}{
		{"1-5", []int{1, 2, 3, 4, 5}, 4, 3},
		{"1-16 (6)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 6, 5},
		{"1-16 (14)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 14, 13},
		{"1-16 (7)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 8, 7},
		{"1-16 (13)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16}, 13, -1},
		{"1-16 (25)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16}, 25, -1},
		{"Пустой массив", []int{}, 25, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BinarySearch(tc.s, tc.i)
			require.Equal(t, tc.expected, got)
		})
	}
}
