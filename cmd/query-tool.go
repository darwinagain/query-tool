package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"query-tool/internal/models"
	"query-tool/internal/postgres"
	"query-tool/internal/stats"

	"github.com/gocarina/gocsv"
)

func main() {
	var file = flag.String("f", "", "path to csv file to use for input")
	var workers = flag.Int("w", 1, "number of concurrent workers")
	flag.Parse()

	log.Printf("file=%v, workers=%d\n", *file, *workers)

	log.Println("reading query parameters...")
	// read input file and convert contents to query parameters objects
	parameters, err := readFile(file)
	if err != nil {
		msg := fmt.Sprintf("error reading csv file: %v", err)
		panic(msg)
	}

	// sort parameters by hostname to ensure all queries for same host go to the same worker
	sortedParameters := sortParameters(parameters)

	// create channels for workers
	numJobs := len(sortedParameters)
	jobs := make(chan []models.QueryParameter, numJobs)
	results := make(chan []models.UsageOutput, numJobs)

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
	var usageOuputs []models.UsageOutput
	for a := 1; a <= numJobs; a++ {
		r := <-results
		usageOuputs = append(usageOuputs, r...)
	}

	// write query results to a csv file
	writeResults(usageOuputs)

	log.Println("calculating benchmark stats now...")
	// get benchmark stats for queries run
	stats, err := stats.GetStats(usageOuputs)
	if err != nil {
		msg := fmt.Sprintf("error calculating stats: %v", err)
		panic(msg)
	}

	// output stats
	log.Printf("number of queries run: %v\n", stats.QueriesRun)
	log.Printf("total processing time: %v\n", stats.TotalTime)
	log.Printf("minimum query processing time: %v\n", stats.MinimumTime)
	log.Printf("median query processing time: %v\n", stats.MedianTime)
	log.Printf("average query processing time: %v\n", stats.AverageTime)
	log.Printf("maxiumum query processing time: %v\n", stats.MaximumTime)
}

func worker(id int, jobs <-chan []models.QueryParameter, results chan<- []models.UsageOutput) {
	for j := range jobs {
		log.Println("worker", id, "started  job for", j[0].HostName)

		// run queries
		var usageOuputs []models.UsageOutput
		for _, v := range j {
			output, err := postgres.RunQuery(v)
			if err != nil {
				log.Printf("error running query for %v: %v", v, err)
			}

			usageOuputs = append(usageOuputs, output)
		}
		results <- usageOuputs
		log.Println("worker", id, "finished job for", j[0].HostName)
	}
}

func sortParameters(parameters []models.QueryParameter) map[string][]models.QueryParameter {
	// sort parameters into map by host
	sortedParameters := make(map[string][]models.QueryParameter)
	for i := range parameters {
		sortedParameters[parameters[i].HostName] = append(sortedParameters[parameters[i].HostName], parameters[i])
	}

	return sortedParameters
}

func readFile(filePath *string) ([]models.QueryParameter, error) {
	// open file
	file, err := os.Open(*filePath)
	if err != nil {
		return []models.QueryParameter{}, err
	}
	defer file.Close()

	// convert file contents to QueryParameter objects
	var queryParameters []models.QueryParameter
	if err := gocsv.UnmarshalFile(file, &queryParameters); err != nil {
		return []models.QueryParameter{}, err
	}

	return queryParameters, nil
}

func writeResults(usageOutputs []models.UsageOutput) error {
	// export usage results to csv file
	file, err := ioutil.TempFile("./", "query_results-*.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	gocsv.MarshalFile(&usageOutputs, file)

	return nil
}
