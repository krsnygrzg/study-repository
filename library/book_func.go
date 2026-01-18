package library

import "time"

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

func (b *Book) Read() {
	ReadTime := time.Now()

	b.Readed = true
	b.ReadTime = &ReadTime
}
