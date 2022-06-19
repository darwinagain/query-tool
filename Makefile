MIGRATIONS_PATH = migrations
PSQL_DATABASE_NAME = homework
PSQL_URL ?= postgres://postgres:pgpass@localhost:5432/$(PSQL_DATABASE_NAME)?sslmode=disable

db_migrate:
	migrate -path=$(MIGRATIONS_PATH) -database=$(PSQL_URL) up
db_migrate_down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(PSQL_URL) down
db_create:
	docker-compose exec database psql -U postgres --command='create database "$(PSQL_DATABASE_NAME)";'
db_setup:
	docker-compose exec -T database psql -U postgres < cpu_usage.sql
db_populate:
	psql $(PSQL_URL) -c "\COPY cpu_usage FROM './migrations/cpu_usage.csv' CSV HEADER"

run:
	docker-compose up -d && \
	psql $(PSQL_URL) -c "\COPY cpu_usage FROM './seeding/cpu_usage.csv' CSV HEADER"
