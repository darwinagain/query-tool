package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
)

type Row struct {
	Timestamp string  `csv:"ts"`
	Host      string  `csv:"host"`
	Usage     float64 `csv:"usage"`
}

func main() {
	log.Println("reading csv file now...")
	rows, err := readFile("./migrations/cpu_usage.csv")
	if err != nil {
		log.Panic(err)
	}

	log.Println("inserting cpu_usage data now...")
	insertUsage(rows)

}

func readFile(filePath string) ([]Row, error) {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		return []Row{}, err
	}
	defer file.Close()

	// convert file contents to Row objects
	var rows []Row
	if err := gocsv.UnmarshalFile(file, &rows); err != nil {
		return []Row{}, err
	}

	return rows, nil
}

func insertUsage(rows []Row) error {
	db := openConnection()
	defer db.Close()
	for _, row := range rows {
		log.Printf("inserting %v", row)
		query := "INSERT INTO cpu_usage (ts, host, usage) VALUES ($1, $2, $3)"
		_, err := db.Exec(query, row.Timestamp, row.Host, row.Usage)
		if err != nil {
			return err
		}
	}

	log.Panicln("successfully inserted all rows into cpu_usage table!")
	return nil

}

func openConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:pgpass@localhost:5432/homework?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
