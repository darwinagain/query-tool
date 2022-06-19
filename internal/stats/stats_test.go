package stats

import (
	"query-tool/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetStats(t *testing.T) {
	expectedStats := models.BenchmarkStats{
		QueriesRun:  2,
		TotalTime:   time.Duration(9000000),
		MinimumTime: time.Duration(3000000),
		MedianTime:  time.Duration(4500000),
		AverageTime: time.Duration(4500000),
		MaximumTime: time.Duration(6000000),
	}

	usageOutputs1 := []models.UsageOutput{
		{
			HostName: "host_000008",
			Minute:   "2017-01-01T08:59:00Z",
			MaxUsage: "51.01",
			MinUsage: "27.54",
		},
		{
			HostName: "host_000008",
			Minute:   "2017-01-01T09:00:00Z",
			MaxUsage: "94.43",
			MinUsage: "2.74",
		},
	}
	usageOutputs2 := []models.UsageOutput{
		{
			HostName: "host_000001",
			Minute:   "2017-01-01T08:59:00Z",
			MaxUsage: "98.39",
			MinUsage: "27.54",
		},
		{
			HostName: "host_000001",
			Minute:   "2017-01-01T09:00:00Z",
			MaxUsage: "98.39",
			MinUsage: "13.26",
		},
	}
	queryResults := []models.QueryResults{
		{
			UsageOutputs: usageOutputs1,
			QueryTime:    time.Duration(6 * time.Millisecond),
		},
		{
			UsageOutputs: usageOutputs2,
			QueryTime:    time.Duration(3 * time.Millisecond),
		},
	}

	stats, err := GetStats(queryResults)
	assert.NoError(t, err)

	assert.Equal(t, expectedStats, stats)
}
