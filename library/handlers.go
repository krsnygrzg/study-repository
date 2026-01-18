package library

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var books = make([]Book, 0)

var m sync.Mutex

func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	var req AddBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	book := NewBook(req.Name, req.Author, req.Pages)

	m.Lock()
	books = append(books, book)
	m.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func HandleGetBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	m.Lock()
	defer m.Unlock()

	for _, book := range books {
		if book.Name == name {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}

func HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	json.NewEncoder(w).Encode(books)
}

func HandleGetReadedBooks(w http.ResponseWriter, r *http.Request) {
	result := make([]Book, 0)

	m.Lock()
	for _, book := range books {
		if book.Readed {
			result = append(result, book)
		}
	}
	m.Unlock()

	json.NewEncoder(w).Encode(result)
}

func HandleGetUnreadedBooks(w http.ResponseWriter, r *http.Request) {
	result := make([]Book, 0)

	m.Lock()
	for _, book := range books {
		if !book.Readed {
			result = append(result, book)
		}
	}
	m.Unlock()

	json.NewEncoder(w).Encode(result)
}

func HandleCompleteBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	m.Lock()
	defer m.Unlock()

	for i := range books {
		if books[i].Name == name {
			books[i].Read()
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}

func HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	m.Lock()
	defer m.Unlock()

	for i := range books {
		if books[i].Name == name {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}
