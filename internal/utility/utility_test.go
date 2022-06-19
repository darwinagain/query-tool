package utility

import (
	"query-tool/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortParameters(t *testing.T) {
	parameters := []models.QueryParameter{
		{
			HostName:  "host_000008",
			StartTime: "2017-01-01 08:59:22",
			EndTime:   "2017-01-01 09:59:22",
		},
		{
			HostName:  "host_000001",
			StartTime: "2017-01-01 08:59:22",
			EndTime:   "2017-01-01 09:59:22",
		},
		{
			HostName:  "host_000008",
			StartTime: "2017-01-02 18:50:28",
			EndTime:   "2017-01-02 19:50:28",
		},
		{
			HostName:  "host_000001",
			StartTime: "2017-01-02 18:50:28",
			EndTime:   "2017-01-02 19:50:28",
		},
		{
			HostName:  "host_000008",
			StartTime: "2017-01-03 18:50:28",
			EndTime:   "2017-01-03 19:50:28",
		},
	}

	sorted := SortParameters(parameters)

	assert.Equal(t, 3, len(sorted["host_000008"]))
	assert.Equal(t, 2, len(sorted["host_000001"]))
}

func TestReadFile(t *testing.T) {
	filePath := "../../query_params.csv"

	parameters, err := ReadFile(&filePath)
	assert.NoError(t, err)

	assert.Equal(t, 200, len(parameters))
}
