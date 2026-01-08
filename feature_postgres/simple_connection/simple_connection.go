package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	} //тестовый запрос в бд

	fmt.Println("Подключение к базе даных было успешно")
}
