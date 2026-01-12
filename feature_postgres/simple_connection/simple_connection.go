package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("conn_String")

	return pgx.Connect(ctx, connString)

}
