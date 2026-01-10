package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// <- http(get book)
// -> db (books)
// rows <- db
// rows-> http

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, name, author, pages, readed, buy_time, read_time
	FROM books
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Pages,
			&book.Readed,
			&book.BuyTime,
			&book.ReadTime,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)

		// printBook(book)
	}

	return books, nil
}

func printBook(book BookModel) {
	fmt.Println("---------------------------")
	fmt.Println("id:", book.ID)
	fmt.Println("name:", book.Name)
	fmt.Println("author:", book.Author)
	fmt.Println("pages:", book.Pages)
	fmt.Println("readed:", book.Readed)
	fmt.Println("buy_time:", book.BuyTime)
	fmt.Println("read_time:", book.ReadTime)
}
