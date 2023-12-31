.PHONY: setup_node build createdb dropdb migrateup migratedown

setup_node:
	cd frontend && npm install

build:
	docker build -t twertle .

postgres:
	docker run --name postgres15 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretbacon -p 5432:5432 -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path databases/migration -database "postgresql://root:secretbacon@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path databases/migration -database "postgresql://root:secretbacon@localhost:5432/simple_bank?sslmode=disable" -verbose down
