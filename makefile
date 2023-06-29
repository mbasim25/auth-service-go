SHELL := /bin/bash

createdb:
	docker-compose exec -it db createdb --username=postgres --owner=postgres void

dropdb:
	docker exec -it db dropdb void

migrate:
		migrate -source file://db/migrations \
			-database postgres://postgres:password@localhost:5435/void?sslmode=disable up

sqlc:
	sqlc generate

test:
	export PATH=/usr/local/go/:$$PATH && go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mbasim25/ticketing-app-microservices/db/sqlc Store

.PHONY: createdb dropdb migrate sqlc test mock
