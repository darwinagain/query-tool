package postgres

import (
	"database/sql"
	"log"
	"query-tool/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestRunQuery(t *testing.T) {
	parameters := models.QueryParameter{
		HostName:  "host_000008",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	}

	db := OpenConnection()

	output, err := RunQuery(db, parameters)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
