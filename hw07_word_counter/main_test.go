package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected map[string]int
	}{
		{"зима", "зима  санки зима Санки санки холод санки", map[string]int{
			"зима":  2,
			"холод": 1,
			"санки": 4,
		}},
		{"животные", "кот собака коты птица  и птица", map[string]int{
			"кот":    1,
			"собака": 1,
			"коты":   1,
			"птица":  2,
			"и":      1,
		}},
		{"пустая строка", "", map[string]int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountWords(tc.s)
			//if got != tc.expected {
			//	t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
			//}
			require.Equal(t, tc.expected, got)
		})
	}
}
