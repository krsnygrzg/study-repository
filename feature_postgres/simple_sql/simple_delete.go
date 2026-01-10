package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, booksID []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, booksID)

	return err
}
