package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
// 	sqlQuery := `
// 	UPDATE books
// 	SET readed = TRUE
// 	WHERE ID = 1;
// 	`

// 	_, err := conn.Exec(ctx, sqlQuery)

// 	return err
// }

func UpdateBook(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) error {
	sqlQuery := `
UPDATE books
SET name = $1 , author = $2, pages = $3, readed = $4, buy_time = $5, read_time = $6
WHERE id=$7;
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
		book.ID,
	)

	return err
}
