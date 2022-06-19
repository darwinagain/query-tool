BUILD_DIR=bin
TEST_COVERAGE_OUT=$(BUILD_DIR)/coverage.out
TEST_FLAGS=-p 1 -count 1 -coverpkg=./... -coverprofile=$(TEST_COVERAGE_OUT)
TEST_PKGS=$(shell go list ./...)

MIGRATIONS_PATH = migrations
PSQL_DATABASE_NAME = homework
PSQL_URL ?= postgres://postgres:pgpass@localhost:5432/$(PSQL_DATABASE_NAME)?sslmode=disable

# these are old and may not work
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


# use these
test:
	mkdir -p $(BUILD_DIR)
	go test $(TEST_FLAGS) $(TEST_PKGS)

run:
	docker-compose up -d && \
	psql $(PSQL_URL) -c "\COPY cpu_usage FROM './seeding/cpu_usage.csv' CSV HEADER"
