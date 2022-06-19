package utility

import (
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
