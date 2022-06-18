# timescale-query-tool

A command line tool to benchmark query performance across multiple workers against a TimescaleDB instance.

## Installation

- [Golang 1.18](https://golang.org/doc/install)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [golang-migrate](https://github.com/golang-migrate/migrate) - `go get -tags 'postgres' -u https://github.com/golang-migrate/migrate`

#### Setting Up Your Development Environment

# Setting Up Your Local Environment

Create docker container for dependencies:

```
docker-compose up -d
```

Run `make db_setup` to set up the database, timescale extension, and table

Run `psql postgres://postgres:pgpass@localhost:5432/homework?sslmode=disable -c "\COPY cpu_usage FROM cpu_usage.csv CSV HEADER"` to populate the `cpu_usage` table

# Usage

To see all available flags run:

```
go run cmd/query-tool.go --help
```

Using the existing file `query_params.csv` as input and your desired number of workers run:

```
go run cmd/query-tool.go -f query_params.csv -w 4
```

## Testing

`Coming Soon`
