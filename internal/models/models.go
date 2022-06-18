package models

import "time"

type QueryParameter struct {
	HostName  string `csv:"hostname"`
	StartTime string `csv:"start_time"`
	EndTime   string `csv:"end_time"`
}

type UsageOutput struct {
	HostName  string        `csv:"hostname"`
	Minute    string        `csv:"minute"`
	MaxUsage  string        `csv:"max_usage"`
	MinUsage  string        `csv:"min_usage"`
	QueryTime time.Duration `csv:"query_time"`
}

type BenchmarkStats struct {
	QueriesRun  int
	TotalTime   time.Duration
	MinimumTime time.Duration
	MedianTime  time.Duration
	AverageTime time.Duration
	MaximumTime time.Duration
}
