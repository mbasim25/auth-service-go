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
	/usr/local/go/bin/go test -v -cover ./...

.PHONY: createdb
