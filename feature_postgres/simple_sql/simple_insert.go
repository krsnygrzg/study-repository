package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) error {
	sqlQuery := `
	INSERT INTO books (name, author, pages, readed, buy_time, read_time)
	VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.Pages,
		book.Readed,
		book.BuyTime,
		book.ReadTime,
	)

	return err
}
