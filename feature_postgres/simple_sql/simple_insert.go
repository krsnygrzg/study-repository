package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	name string,
	author string,
	pages int,
) error {
	sqlQuery := `
	INSERT INTO books (name, author, pages)
	VALUES ($1, $2, $3);
	`
	_, err := conn.Exec(ctx, sqlQuery, name, author, pages)

	return err
}
