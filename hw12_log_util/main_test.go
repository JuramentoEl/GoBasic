package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultData(t *testing.T) {
	tests := []struct {
		name               string
		pathFile           string
		levelLog           string
		pathOutput         string
		pathFileExpected   string
		levelLogExpected   string
		pathOutputExpected string
	}{
		{"Пусто", "", "", "", "log.txt", "info", "statistics.txt"},
		{"Заполнено", "statistics.txt", "error", "log.txt", "statistics.txt", "error", "log.txt"},
		{"Заполнено частично", "", "error", "", "log.txt", "error", "statistics.txt"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := defaultData(&tc.pathFile, &tc.levelLog, &tc.pathOutput)
			require.Equal(t, err, nil)
			require.Equal(t, tc.pathFile, tc.pathFileExpected)
			require.Equal(t, tc.levelLog, tc.levelLogExpected)
			require.Equal(t, tc.pathOutput, tc.pathOutputExpected)
		})
	}
}

func TestReadCsv(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		data     map[string]int
		expected []string
	}{
		{"test", 15, map[string]int{"error": 2}, []string{"error: 2", "Количество записей: 15"}},
		{"test", 3, map[string]int{}, []string{"Количество записей: 3"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var linesSlice []string
			writeCsv("text.txt", tc.count, tc.data)
			lines, err := readCsv("text.txt")
			for _, line := range lines {
				linesSlice = append(linesSlice, line[0])
			}
			require.Equal(t, err, nil)
			require.Equal(t, linesSlice, tc.expected)
		})
	}
}
