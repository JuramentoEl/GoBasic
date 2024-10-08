package main

import (
	"fmt"
)

type IndexField int

type Comparator struct {
	fieldCompare IndexField
}

func (c *Comparator) FieldCompare() IndexField {
	return c.fieldCompare
}

func (c *Comparator) SetFieldCompare(fieldCompare IndexField) {
	c.fieldCompare = fieldCompare
}

func NewComparator(fieldComapre IndexField) *Comparator {
	comparator := Comparator{
		fieldCompare: fieldComapre,
	}

	return &comparator
}

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

const (
	Year IndexField = iota + 1
	Size
	Rate
)

func (b *Book) getBook() bool {
	fmt.Println("ID: ", b.ID())
	fmt.Println("Title: ", b.Title())
	fmt.Println("Author: ", b.Author())
	fmt.Println("Year: ", b.Year())
	fmt.Println("Size: ", b.Size())
	fmt.Println("Rate: ", b.Rate())

	return true
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

func (c Comparator) Compare(book1, book2 *Book) bool {
	switch c.FieldCompare() {
	case Year:
		return book1.Year() > book2.Year()
	case Size:
		return book1.Size() > book2.Size()
	case Rate:
		return book1.Rate() > book2.Rate()
	}
	return false
}

func main() {
	books := make(map[int]*Book)
	count := 0

	actionSelection(&count, books)
}

func actionSelection(count *int, books map[int]*Book) {
	var action int

	fmt.Println("1: сохранить книгу")
	fmt.Println("2: получить данные о книге")
	fmt.Println("3: сравнить книги")
	fmt.Print("Выберете действие: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		*count++
		books[*count] = scanBook(*count)
		actionSelection(count, books)
	case 2:
		var IDBook int
		fmt.Print("В базе ", *count, " книг. Выберете одну: ")
		fmt.Scanln(&IDBook)
		book := books[IDBook]
		book.getBook()
		actionSelection(count, books)
	case 3:
		var bookIndex1, bookIndex2 int
		var filed IndexField
		fmt.Print("В базе ", *count, " книг. ")
		fmt.Println("Введите номера книг для сранение и имя поля через пробел (1 - year, 2 - size или  3 - rate)")
		fmt.Scanln(&bookIndex1, &bookIndex2, &filed)
		comparator := NewComparator(filed)
		result := comparator.Compare(books[bookIndex1], books[bookIndex2])
		fmt.Println(result)
		actionSelection(count, books)
	}
}

func scanBook(id int) *Book {
	var title, author string
	var year, size int
	var rate float32

	fmt.Println("Ввод данных книги")
	fmt.Print("Название: ")
	fmt.Scanln(&title)
	fmt.Print("Автор: ")
	fmt.Scanln(&author)
	fmt.Print("Год: ")
	fmt.Scanln(&year)
	fmt.Print("Размер: ")
	fmt.Scanln(&size)
	fmt.Print("Редкость: ")
	fmt.Scanln(&rate)
	book := NewBook(id, title, author, year, size, rate)

	return book
}
