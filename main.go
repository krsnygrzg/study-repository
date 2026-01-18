package main

import (
	"context"
	"log"

	simpleconnection "library/feature_postgres/simple_connection"
	"library/feature_postgres/simple_sql"
	"library/library"
)

func main() {
	ctx := context.Background()

	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	log.Println("Server started on :9091")

	if err := library.StartServer(); err != nil {
		log.Fatal(err)
	}
}
