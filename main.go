package main

import (
	"context"
	"log"
	"time"

	simpleconnection "library/feature_postgres/simple_connection"
	"library/feature_postgres/simple_sql"
	"library/library"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	log.Println("server started on :9091")
	if err := library.StartServer(conn, "9091"); err != nil {
		log.Fatal(err)
	}
}
