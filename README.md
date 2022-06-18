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

You will see log messages while the csv file is read and the queries are run and an eventual output similar to the following:

```
2022/06/18 15:03:40 number of queries run: 200
2022/06/18 15:03:40 total processing time: 773.220341ms
2022/06/18 15:03:40 minimum query processing time: 3.507686ms
2022/06/18 15:03:40 median query processing time: 3.813921ms
2022/06/18 15:03:40 average query processing time: 3.866101ms
2022/06/18 15:03:40 maxiumum query processing time: 5.081698ms
```

### Testing

`Coming Soon`
