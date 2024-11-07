package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArithmeticMean(t *testing.T) {
	tests := []struct {
		name     string
		sensor   []int64
		expected []int64
	}{
		{"1", []int64{79, 85, 50, 98, 2, 93, 39, 37, 31, 3}, []int64{51}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sensorChannel := make(chan int64)
			meanChannel := make(chan int64)
			go arithmeticMean(sensorChannel, meanChannel)
			var result []int64
			for _, sensor := range tc.sensor {
				sensorChannel <- sensor
			}
			close(sensorChannel)
			for mean := range meanChannel {
				result = append(result, mean)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}
