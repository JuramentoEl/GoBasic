package main

import (
	"testing"
)

func TestGetChessboerd(t *testing.T) {
	const size, want = 3, " # \r\n# #\r\n # "
	got := getChessboerd(3)
	if got != want {
		t.Errorf("getChessboerd(%q) = %v; want %v", size, got, want)
	}
}
