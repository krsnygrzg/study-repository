include .env
export

service-run: 
	go run main.go
migrate-up:
	migrate -path feature_postgres/migrations -database ${CONN_STRING} up 
migrate-down:
	migrate -path feature_postgres/migrations -database ${CONN_STRING} down  