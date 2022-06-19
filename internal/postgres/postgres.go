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

func RunQuery(db *sql.DB, v models.QueryParameter) (models.QueryResults, error) {
	query := `select host, date_trunc('minute', ts) AS MINUTE, max(usage) as max_cpu, min(usage) as min_cpu
			from cpu_usage
			where
				host= $1
				and ts >= $2
				and ts <= $3
			group by host, minute;`

	var output []models.UsageOutput
	start := time.Now()
	rows, err := db.Query(query, v.HostName, v.StartTime, v.EndTime)
	if err != nil {
		return models.QueryResults{}, err
	}
	defer rows.Close()
	for rows.Next() {
		usageOutput := new(models.UsageOutput)
		if err := rows.Scan(&usageOutput.HostName, &usageOutput.Minute, &usageOutput.MaxUsage, &usageOutput.MinUsage); err != nil {
			return models.QueryResults{}, err
		}
		output = append(output, *usageOutput)
	}

	elapsed := time.Since(start)
	results := models.QueryResults{
		UsageOutputs: output,
		QueryTime:    elapsed,
	}

	return results, nil
}
