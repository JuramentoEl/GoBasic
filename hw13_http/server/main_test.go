package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAddr(t *testing.T) {
	tests := []struct {
		server string
		port   string
		result string
	}{
		{"localhost", "8080", "localhost:8080"},
		{"localhost", "8000", "localhost:8000"},
		{"87.117.35.71", "8080", "87.117.35.71:8080"},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			fullURL := getAddr(tc.server, tc.port)
			require.Equal(t, fullURL, tc.result)
		})
	}
}
