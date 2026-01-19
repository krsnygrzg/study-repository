package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) (BookModel, error) {

	sqlQuery := `
	INSERT INTO books (name, author, pages, readed, buy_time, read_time)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING name, author, pages, readed, buy_time, read_time;
	`

	var result BookModel

	err := conn.QueryRow(
		ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.Pages,
		book.Readed,
		book.BuyTime,
		book.ReadTime,
	).Scan(
		&result.Name,
		&result.Author,
		&result.Pages,
		&result.Readed,
		&result.BuyTime,
		&result.ReadTime,
	)

	return result, err
}
