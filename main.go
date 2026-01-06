package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	Name     string     `json:"name"`
	Author   string     `json:"author"`
	Pages    int        `json:"pages"`
	Readed   bool       `json:"readed"`
	BuyTime  time.Time  `json:"buy_time"`
	ReadTime *time.Time `json:"read_time"`
}

type AddBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func NewBook(name string, author string, pages int) Book {
	return Book{
		Name:    name,
		Author:  author,
		Pages:   pages,
		Readed:  false,
		BuyTime: time.Now(),
	}
}

// func (b *Book) Read() {
// 	ReadTime := time.Now()

// 	b.Readed = true
// 	b.ReadTime = &ReadTime
// }

var books = make([]Book, 0)

var m sync.Mutex

func StartServer() error {
	router := mux.NewRouter()
	router.Path("/book").Methods("POST").HandlerFunc(HandleAddBook)
	router.Path("/book/{name}").Methods("GET").HandlerFunc(HandleGetBook)
	// router.Path("/book").Methods("GET").HandlerFunc(HandleGetAllBooks)
	// router.Path("/book/{name}").Methods("GET").Queries("readed", "true").HandlerFunc(HandleGetReadedBooks)
	// router.Path("/book/{name}").Methods("GET").Queries("readed", "false").HandlerFunc(HandleGetUnreadedBooks)
	// router.Path("/book/{name}").Methods("PATCH").HandlerFunc(HandleCompleteBook)
	// router.Path("/book/{name}").Methods("DELETE").HandlerFunc(HandleDeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}

func HandleAddBook(w http.ResponseWriter, r *http.Request) {

	var req AddBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	book := NewBook(req.Name, req.Author, req.Pages)

	m.Lock()
	defer m.Unlock()

	books = append(books, book)

	b, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}

func HandleGetBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	m.Lock()
	defer m.Unlock()

	for _, book := range books {
		if book.Name == name {
			b, err := json.Marshal(book)
			if err != nil {
				http.Error(w, "failed to marshal book", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}

	}
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal(err)
	}
}
