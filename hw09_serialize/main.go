package main

import (
	"encoding/json"
	"fmt"

	"github.com/JuramentoEl/GoBasic/hw09_serialize/ProtoBook"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

func (b *Book) GetID() int {
	return b.ID
}

func (b *Book) DetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetYear() int {
	return b.Year
}

func (b *Book) GetSize() int {
	return b.Size
}

func (b *Book) GateRate() float32 {
	return b.Rate
}

func (b *Book) SetID(id int) {
	b.ID = id
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (b *Book) SetYear(year int) {
	b.Year = year
}

func (b *Book) SetSize(size int) {
	b.Size = size
}

func (b *Book) SetRate(rate float32) {
	b.Rate = rate
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

func main() {
	book1 := NewBook(1, "Мастер и Маргарита", "Михаил Булгаков", 2018, 512, 2)
	JSONTransformation(book1)

	book2 := &ProtoBook.Book{
		Id:     2,
		Title:  "Преступление и наказание",
		Author: "Федор Достоевский",
		Year:   2023,
		Size:   592,
		Rate:   6,
	}
	ProtoTransformation(book2)
}

func JSONTransformation(book *Book) {
	j, err := JSONMarshal(book)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("book json: %s\n", j)
	book2, err := JSONUnmarshal(j)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("book struct: %v\n", *book2)
}

func JSONMarshal(v *Book) ([]byte, error) {
	j, err := json.Marshal(v)
	return j, err
}

func JSONUnmarshal(j []byte) (*Book, error) {
	var book Book
	err := json.Unmarshal(j, &book)
	return &book, err
}

func ProtoTransformation(book *ProtoBook.Book) {
	var book2 ProtoBook.Book
	data, err := ProtoMarshal(book)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("book proto: %s\n", data)

	err = proto.Unmarshal(data, &book2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("book message: %v\n", &book2)
}

func ProtoMarshal(book *ProtoBook.Book) ([]byte, error) {
	data, err := proto.Marshal(book)
	return data, err
}

func ProtoUnmarshal(j []byte) (*ProtoBook.Book, error) {
	var book ProtoBook.Book

	err := proto.Unmarshal(j, &book)
	return &book, err
}
