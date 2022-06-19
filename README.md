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

**Note: Some tests require database access to run, so ensure the database is spun up in docker via `make run`**

You can run all unit tests from the root directory of the project with the following command:

```
make test
```

### Assumptions

1. The csv file used for input will always be in the format present in `query_params.csv` (hostname, start_time, end_time)
2. The user will specify the full filepath for the input csv file
3. The database will always be populated with the provided cpu_usage.csv file present in `/seeding`

### With More Time I Would...

1. Make unit tests more robust to handle error and edge cases
2. Add more benchmark statistics
3. Export the results from all queries to a csv file or database table for historical inspection
