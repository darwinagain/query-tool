# timescale-query-tool

A command line tool to benchmark query performance across multiple workers against a TimescaleDB instance.

## Installation

- [Golang 1.16+](https://golang.org/doc/install)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [golang-migrate](https://github.com/golang-migrate/migrate) - `go get -u -d github.com/golang-migrate/migrate/cmd/migrate`

### Setting Up Your Development Environment

Create docker container for dependencies:

```
docker-compose up -d
```

Run `make db_setup` to set up the database, timescale extension, and table

Run `psql postgres://postgres:pgpass@localhost:5432/homework?sslmode=disable -c "\COPY cpu_usage FROM cpu_usage.csv CSV HEADER"` to populate the `cpu_usage` table

# Usage

### Running the Tool Locally

To see all available flags run the following command:

```
go run cmd/query-tool.go --help
```

Using the existing csv file in this project `query_params.csv` as input and your desired number of workers run the following command:

```
go run cmd/query-tool.go -f query_params.csv -w 4
```

### Testing

`Coming Soon`
