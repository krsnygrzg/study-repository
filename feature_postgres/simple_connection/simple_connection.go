package simpleconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")

}
