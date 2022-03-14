mkdir -p db/migration
migrate create -ext sql -dir db/migration -seq init_schema
migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" -verbose up