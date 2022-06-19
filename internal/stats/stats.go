package stats

import (
	"query-tool/internal/models"
	"sort"
	"time"
)

func GetStats(queryResults []models.QueryResults) (models.BenchmarkStats, error) {
	var times []time.Duration
	// create list of query times
	for _, v := range queryResults {
		times = append(times, v.QueryTime)
	}

	// sort query times
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})

	// get all stats
	stats := models.BenchmarkStats{
		QueriesRun:  len(times),
		TotalTime:   sum(times),
		MinimumTime: times[0],
		MedianTime:  median(times),
		AverageTime: sum(times) / time.Duration(len(times)),
		MaximumTime: times[len(times)-1],
	}

	return stats, nil
}

func sum(arr []time.Duration) time.Duration {
	var result time.Duration
	// add durations together
	for _, i := range arr {
		result += i
	}

	return result
}

func median(times []time.Duration) time.Duration {
	var median time.Duration
	l := len(times)
	if l == 0 {
		// length is 0
		return 0
	} else if l%2 == 0 {
		// length is even so add together two middle value
		median = (times[l/2-1] + times[l/2]) / 2
	} else {
		// length is odd so take middle value
		median = times[l/2]
	}

	return median
}
