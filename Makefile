postgres:
	docker run --name=postgres16 -e POSTGRES_USER=root -p 5433:5432 -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

.PHONEY: postgres createdb dropdb