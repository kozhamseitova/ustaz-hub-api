.PHONY:
migrate-up:
	migrate -source file://./migrations/pg -database postgres://postgres:password@localhost:5432/ustaz_hub_db?sslmode=disable up

.PHONY:
create-migration:
	migrate create -ext sql -dir schema -seq ${name}