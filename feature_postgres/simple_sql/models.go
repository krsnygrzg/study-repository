package simple_sql

import "time"

type BookModel struct {
	ID       int
	Name     string
	Author   string
	Pages    int
	Readed   bool
	BuyTime  time.Time
	ReadTime *time.Time
}
