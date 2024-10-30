package main

import (
	"testing"

	"github.com/JuramentoEl/GoBasic/hw09_serialize/ProtoBook"
	"github.com/stretchr/testify/require"
)

func TestJSONMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		book *Book
	}{
		{"Книга 1", NewBook(1, "Мастер и Маргарита", "Михаил Булгаков", 2018, 512, 2)},
		{"Книга 2", NewBook(2, "Преступление и наказание", "Федор Достоевский", 2023, 592, 6)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			j, _ := JSONMarshal(tc.book)
			st, _ := JSONUnmarshal(j)
			require.Equal(t, tc.book, st)
		})
	}
}

func TestProtoMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		book *ProtoBook.Book
	}{
		{"Книга 1", &ProtoBook.Book{
			Id:     1,
			Title:  "Мастер и Маргарита",
			Author: "Михаил Булгаков",
			Year:   2018,
			Size:   512,
			Rate:   2,
		}},
		{"Книга 2", &ProtoBook.Book{
			Id:     2,
			Title:  "Преступление и наказание",
			Author: "Федор Достоевский",
			Year:   2023,
			Size:   592,
			Rate:   6,
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			j, _ := ProtoMarshal(tc.book)
			st, _ := ProtoUnmarshal(j)
			require.Equal(t, tc.book.String(), st.String())
		})
	}
}
