package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// func ResetDB(ctx context.Context, conn *pgx.Conn) error {
// 	_, err := conn.Exec(ctx, `
// 		DROP TABLE IF EXISTS books;
// 	`)
// 	return err
// }

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
         id SERIAL PRIMARY KEY,
		 name VARCHAR(100) NOT NULL,
		 author VARCHAR(100) NOT NULL,
		 pages INTEGER NOT NULL CHECK(pages <= 10000),
		 readed BOOLEAN NOT NULL DEFAULT false,
		 buy_time TIMESTAMPTZ NOT NULL DEFAULT now(),
		 read_time TIMESTAMPTZ,
		 UNIQUE(name)
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
