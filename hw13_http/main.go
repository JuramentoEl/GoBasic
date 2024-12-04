package main

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBook(id int, title, author string, year, size int, rate float32) *Book {
	book := Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}

	return &book
}

func main() {
}
