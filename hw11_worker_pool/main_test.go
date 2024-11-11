package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected int
	}{
		{"500", 500, 500},
		{"1000", 1000, 1000},
		{"0", 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := counter(tc.count)
			require.Equal(t, tc.expected, got)
		})
	}
}
