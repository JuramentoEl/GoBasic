package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUrl(t *testing.T) {
	tests := []struct {
		server string
		port   string
		url    string
		result string
	}{
		{"localhost", "8080", "v1/getBook", "http://localhost:8080/v1/getBook"},
		{"localhost", "8000", "v1/getPerson", "http://localhost:8000/v1/getPerson"},
		{"87.117.35.71", "8080", "v1/createBook", "http://87.117.35.71:8080/v1/createBook"},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			fullURL := getURL(tc.server, tc.port, tc.url)
			require.Equal(t, fullURL, tc.result)
		})
	}
}
