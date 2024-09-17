package main

import (
	"fmt"
)

type book struct {
	ID              int
	Title           string
	Author          string
	Year            int
	Size            int
	Rate            float32
	FiledComparison indexField
}

type indexField int

const (
	Year indexField = iota + 1
	Size
	Rate
)

func (b *book) setBook(id int, title string, author string, year int, size int, rate float32) book {
	book := book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
		Size:   size,
		Rate:   rate,
	}

	return book
}

func (b *book) getBook() bool {
	fmt.Println("ID: ", b.ID)
	fmt.Println("Title: ", b.Title)
	fmt.Println("Author: ", b.Author)
	fmt.Println("Year: ", b.Year)
	fmt.Println("Size: ", b.Size)
	fmt.Println("Rate: ", b.Rate)

	return true
}

func (b *book) setFiledComparison(FiledComparison indexField) {
	b.FiledComparison = FiledComparison
}

func (b *book) getFiledComparison() indexField {
	return b.FiledComparison
}

func main() {
	books := make(map[int]book)
	count := 0

	actionSelection(&count, books)
}

func actionSelection(count *int, books map[int]book) {
	var action int
	var book2, book1 book

	fmt.Println("1: сохранить книгу")
	fmt.Println("2: получить данные о книге")
	fmt.Println("3: сравнить книги")
	fmt.Print("Выберете действие: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		*count++
		books[*count] = scanBook(*count)
		book2 = books[*count]
		actionSelection(count, books)
	case 2:
		var IDBook int
		fmt.Print("В базе ", *count, " книг. Выберете одну: ")
		fmt.Scanln(&IDBook)
		book2 = books[IDBook]
		book2.getBook()
		actionSelection(count, books)
	case 3:
		var bookIndex1, bookIndex2 int
		var filed indexField
		var result bool
		fmt.Print("В базе ", *count, " книг. ")
		fmt.Println("Введите номера книг для сранение и имя поля через пробел (1 - year, 2 - size или  3 - rate)")
		fmt.Scanln(&bookIndex1, &bookIndex2, &filed)
		book2 = books[bookIndex1]
		book2.setFiledComparison(filed)
		book1 = books[bookIndex2]
		book1.setFiledComparison(filed)
		result = comparisonBooks(book2, book1)
		fmt.Println(result)
		actionSelection(count, books)
	}
}

func scanBook(id int) book {
	var title, author string
	var year, size int
	var rate float32
	var book book

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
	book = book.setBook(id, title, author, year, size, rate)

	return book
}

func comparisonBooks(book1 book, book2 book) bool {
	switch book2.getFiledComparison() {
	case Year:
		return book1.Year > book2.Year
	case Size:
		return book1.Size > book2.Size
	case Rate:
		return book1.Rate > book2.Rate
	}
	fmt.Println("book1: ", book1.FiledComparison)
	fmt.Println("book2: ", book2.FiledComparison)
	return false
}
