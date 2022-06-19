package models

import "time"

type QueryParameter struct {
	HostName  string `csv:"hostname"`
	StartTime string `csv:"start_time"`
	EndTime   string `csv:"end_time"`
}

type UsageOutput struct {
	HostName string
	Minute   string
	MaxUsage string
	MinUsage string
}

type QueryResults struct {
	UsageOutputs []UsageOutput
	QueryTime    time.Duration
}

type BenchmarkStats struct {
	QueriesRun  int
	TotalTime   time.Duration
	MinimumTime time.Duration
	MedianTime  time.Duration
	AverageTime time.Duration
	MaximumTime time.Duration
}
