package main

import (
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

type indexField int

const (
	Year indexField = iota + 1
	Size
	Rate
)

func (b *Book) saveBook(ID int, title string, author string, year int, size int, rate float32) Book {
	book := Book{
		ID:     ID,
		Title:  title,
		Author: author,
		Year:   year,
		Size:   size,
		Rate:   rate,
	}

	return book
}

func (b *Book) getBook() bool {
	fmt.Println("ID: ", b.ID)
	fmt.Println("Title: ", b.Title)
	fmt.Println("Author: ", b.Author)
	fmt.Println("Year: ", b.Year)
	fmt.Println("Size: ", b.Size)
	fmt.Println("Rate: ", b.Rate)

	return true
}

func main() {
	books := make(map[int]Book)
	count := 0

	actionSelection(&count, books)
}

func actionSelection(count *int, books map[int]Book) {

	var action int
	var book Book

	fmt.Println("1: сохранить книгу")
	fmt.Println("2: получить данные о книге")
	fmt.Println("3: сравнить книги")
	fmt.Print("Выберете действие: ")
	fmt.Scanln(&action)

	switch action {
	case 1:
		*count = *count + 1
		books[*count] = scanBook(*count)
		book = books[*count]
		actionSelection(count, books)
	case 2:
		var IDBook int
		fmt.Print("В базе ", *count, " книг. Выберете одну: ")
		fmt.Scanln(&IDBook)
		book = books[IDBook]
		book.getBook()
		actionSelection(count, books)
	case 3:
		var book1, book2 int
		var filed indexField
		var result bool
		fmt.Println("В базе ", *count, " книг. Введите номера книг для сранение и имя поля через пробел (1 - year, 2 - size или  3 - rate)")
		fmt.Scanln(&book1, &book2, &filed)
		result = comparisonBooks(books[book1], books[book2], filed)
		fmt.Println(result)
		actionSelection(count, books)
	}
}

func scanBook(ID int) Book {

	var title, author string
	var year, size int
	var rate float32
	var book Book

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

	book = book.saveBook(ID, title, author, year, size, rate)

	return book
}

func comparisonBooks(book1 Book, book2 Book, filed indexField) bool {

	switch filed {
	case 1:
		return book1.Year > book2.Year
	case 2:
		return book1.Size > book2.Size
	case 3:
		return book1.Rate > book2.Rate
	}

	return false

}

//Реализуйте структуру с методом позволяющим сравнивать книги по полям Year, Size, Rate.
//Выбор режима сравнения задается в конструкторе структуры через перечисление (enum).
//Метод принимает 2 книги и выдает true если первый аргумент больше второго и false если наоборот.
//Обратите внимание на value и pointer ресиверы
