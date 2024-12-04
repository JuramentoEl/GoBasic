package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JuramentoEl/GoBasic/hw13_http/types"
)

func main() {
	var server, port, url string

	flag.StringVar(&server, "server", "localhost", "Адрес сервера")
	flag.StringVar(&port, "port", "8080", "Порт")
	flag.StringVar(&url, "url", "v1/getBook", "Путь к ресурсу")
	flag.Parse()

	if url == "v1/creatBook" {
		createBook(server, port, url)
	} else {
		getMethod(server, port, url)
	}
}

func getMethod(server, port, url string) {
	client := &http.Client{}
	fullURL := getURL(server, port, url)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fullURL, nil)
	if err != nil {
		fmt.Println("Ошибка запроса ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ", err)
		return
	}

	fmt.Println(string(body))
}

func createBook(server, port, url string) {
	book := types.NewBook(2, "Война и мир", "Толстой", 1891, 2005, 2)

	client := &http.Client{}
	fullURL := getURL(server, port, url)
	bodyRequest := strings.NewReader(book.String())
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, fullURL, bodyRequest)
	if err != nil {
		fmt.Println("Ошибка запроса ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ", err)
		return
	}

	fmt.Println(string(body))
}

func getURL(server, port, url string) string {
	return "http://" + server + ":" + port + "/" + url
}
