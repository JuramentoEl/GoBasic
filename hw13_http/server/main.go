package main

import (
	"fmt"
	"net/http"

	bookStruct "github.com/JuramentoEl/GoBasic/hw13_http"
)

func main() {
	//http.HandleFunc("/v1/getBook", getBook)
	//http.ListenAndServe(":8080", nil)
	//book := book.NewBook(1, "Мастер и Маргарита", "Булгаков", 1991, 205, 5)
	book := bookStruct.NewBook(1, "Мастер и Маргарита", "Булгаков", 1991, 205, 5)
	fmt.Print(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	//book := book.NewBook(1)
	fmt.Fprint(w, "Hello, World!")
}
