package main

import (
	"testing"
)

func TestGetChessboerd(t *testing.T) {
	const size, want = 3, " # \r\n# #\r\n # "
	got := GetChessboerd(3)
	if got != want {
		t.Errorf("GetChessboerd(%q) = %v; want %v", size, got, want)
	}
}
