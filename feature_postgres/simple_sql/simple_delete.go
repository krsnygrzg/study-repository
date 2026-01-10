package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = 3;
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
