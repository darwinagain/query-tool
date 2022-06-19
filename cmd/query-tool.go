package main

import (
	"flag"
	"fmt"
	"log"
	"query-tool/internal/models"
	"query-tool/internal/postgres"
	"query-tool/internal/stats"
	"query-tool/internal/utility"
)

func main() {
	var file = flag.String("f", "", "path to csv file to use for input with query parameters (hostname, start_time, end_time)")
	var workers = flag.Int("w", 1, "number of concurrent workers to use when running queries")
	flag.Parse()

	log.Printf("file=%v, workers=%d\n", *file, *workers)

	stats := Run(file, workers)

	// output stats
	log.Printf("number of queries run: %v\n", stats.QueriesRun)
	log.Printf("total processing time: %v\n", stats.TotalTime)
	log.Printf("minimum query processing time: %v\n", stats.MinimumTime)
	log.Printf("median query processing time: %v\n", stats.MedianTime)
	log.Printf("average query processing time: %v\n", stats.AverageTime)
	log.Printf("maxiumum query processing time: %v\n", stats.MaximumTime)
}

func Run(filePath *string, workers *int) models.BenchmarkStats {
	log.Println("reading query parameters...")
	// check headers
	err := utility.CheckHeaders(filePath)
	if err != nil {
		msg := fmt.Sprintf("error reading csv file: %v", err)
		log.Fatal(msg)
	}

	// read input file and convert contents to query parameters objects
	parameters, err := utility.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("error reading csv file: %v", err)
		log.Fatal(msg)
	}

	// sort parameters by hostname to ensure all queries for same host go to the same worker
	sortedParameters := utility.SortParameters(parameters)

	// create channels for workers
	numJobs := len(sortedParameters)
	jobs := make(chan []models.QueryParameter, numJobs)
	results := make(chan []models.QueryResults, numJobs)

	// set up workers
	for w := 1; w <= *workers; w++ {
		go worker(w, jobs, results)
	}

	log.Println("running queries now...")
	// send jobs to workers
	for _, v := range sortedParameters {
		jobs <- v
	}
	close(jobs)

	// get results from all workers
	var allResults []models.QueryResults
	for a := 1; a <= numJobs; a++ {
		r := <-results
		allResults = append(allResults, r...)
	}

	log.Println("calculating benchmark stats now...")
	// get benchmark stats for queries run
	stats, err := stats.GetStats(allResults)
	if err != nil {
		msg := fmt.Sprintf("error calculating stats: %v", err)
		panic(msg)
	}

	return stats
}

func worker(id int, jobs <-chan []models.QueryParameter, results chan<- []models.QueryResults) {
	for j := range jobs {
		// connect to database
		db := postgres.OpenConnection()
		defer db.Close()
		log.Println("worker", id, "started  job for", j[0].HostName)

		// run queries
		var usageOuputs []models.QueryResults
		for _, v := range j {
			output, err := postgres.RunQuery(db, v)
			if err != nil {
				log.Printf("error running query for %v: %v", v, err)
			}

			usageOuputs = append(usageOuputs, output)
		}
		results <- usageOuputs
		log.Println("worker", id, "finished job for", j[0].HostName)
	}
}
