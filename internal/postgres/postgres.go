package postgres

import (
	"database/sql"

	"query-tool/internal/models"
	"time"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
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

func RunQuery(v models.QueryParameter) (models.UsageOutput, error) {
	query := `select host, date_trunc('minute', ts) AS MINUTE, max(usage) as max_cpu, min(usage) as min_cpu
			from cpu_usage
			where
				host= $1
				and ts >= $2
				and ts <= $3
			group by host, minute;`
	// connect to database
	db := OpenConnection()
	defer db.Close()

	var output models.UsageOutput
	start := time.Now()
	err := db.QueryRow(query, v.HostName, v.StartTime, v.EndTime).Scan(&output.HostName, &output.Minute, &output.MaxUsage, &output.MinUsage)
	if err != nil {
		return models.UsageOutput{}, err
	}
	elapsed := time.Since(start)
	output.QueryTime = elapsed

	return output, nil
}
