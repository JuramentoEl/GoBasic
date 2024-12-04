package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/JuramentoEl/GoBasic/hw13_http/types"
)

func main() {
	var addServer, port string

	flag.StringVar(&addServer, "server", "localhost", "Адрес сервера")
	flag.StringVar(&port, "port", "8080", "Порт")
	flag.Parse()

	http.HandleFunc("/v1/getBook", getBook)
	http.HandleFunc("/v1/creatBook", creatBook)
	addr := getAddr(addServer, port)
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "appl ication/json")
	book := types.NewBook(1, "Мастер и Маргарита", "Булгаков", 1991, 205, 5)
	json.NewEncoder(w).Encode(*book)
}

func creatBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newBook types.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON : %v", err)
		return
	}

	fmt.Printf("New book: %+v\n", newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func getAddr(server, port string) string {
	return server + ":" + port
}
