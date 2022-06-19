# timescale-query-tool

A command line tool to benchmark query performance across multiple workers against a TimescaleDB instance.

## Installation

You may need to install these if you do not already have them:

- [Golang 1.16+](https://golang.org/doc/install)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)

### Setting Up Your Development Environment

To begin, create docker containers for dependencies and populate the `cpu_usage` table with data by running the following command:

```
make run
```

## Usage

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
