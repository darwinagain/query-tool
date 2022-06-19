package utility

import (
	"encoding/csv"
	"errors"
	"os"
	"query-tool/internal/models"

	"github.com/gocarina/gocsv"
)

func SortParameters(parameters []models.QueryParameter) map[string][]models.QueryParameter {
	// sort parameters into map by host
	sortedParameters := make(map[string][]models.QueryParameter)
	for i := range parameters {
		sortedParameters[parameters[i].HostName] = append(sortedParameters[parameters[i].HostName], parameters[i])
	}

	return sortedParameters
}

func ReadFile(filePath *string) ([]models.QueryParameter, error) {
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

func CheckHeaders(filePath *string) error {
	file, err := os.Open(*filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	i := 0
	for _, c := range headers {
		if c == "hostname" {
			i += 1
		}
		if c == "start_time" {
			i += 1
		}
		if c == "end_time" {
			i += 1
		}
	}

	if i != 3 {
		return errors.New("missing required headers (hostname, start_time, end_time)")
	}

	return nil

}
