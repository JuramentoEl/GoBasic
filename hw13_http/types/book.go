package types

import (
	"encoding/json"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float32 `json:"rate"`
}

func NewBook(id int, title, author string, year, size int, rate float32) *Book {
	book := Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
		Size:   size,
		Rate:   rate,
	}

	return &book
}

func (b Book) String() string {
	body, _ := json.Marshal(b)
	return string(body)
}
