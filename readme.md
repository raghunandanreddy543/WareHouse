mkdir -p db/migration
migrate create -ext sql -dir db/migration -seq init_schema
migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" -verbose up
docker run -d --name postgres postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432
docker run --name posttest -d -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword postgres