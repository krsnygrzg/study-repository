package library

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/book", HandleAddBook).Methods("POST")
	router.HandleFunc("/book", HandleGetAllBooks).Methods("GET")
	router.HandleFunc("/book", HandleGetReadedBooks).Methods("GET").Queries("readed", "true")
	router.HandleFunc("/book", HandleGetUnreadedBooks).Methods("GET").Queries("readed", "false")
	router.HandleFunc("/book/{name}", HandleGetBook).Methods("GET")
	router.HandleFunc("/book/{name}", HandleCompleteBook).Methods("PATCH")
	router.HandleFunc("/book/{name}", HandleDeleteBook).Methods("DELETE")

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
