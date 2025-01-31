package data

import (
	"encoding/csv"
	"os"
	"strconv"
)

// load data from test_data.csv
var Data []map[string]string

func LoadData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	var records []map[string]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		row := make(map[string]string)
		for i, header := range headers {
			row[header] = record[i]
		}
		records = append(records, row)
	}
	Data = records
	return nil
}

func ConsolidateRecords(columns []string, volumes string) []map[string]string {
	consolidated := make(map[string]map[string]string)
	for _, record := range Data {
		key := ""
		for _, column := range columns {
			key += record[column] + "|"
		}
		if _, exists := consolidated[key]; !exists {
			consolidated[key] = make(map[string]string)
			for k, v := range record {
				consolidated[key][k] = v
			}
		} else {
			value, err := strconv.ParseFloat(record[volumes], 64)
			if err != nil {
				continue
			}
			existingValue, err := strconv.ParseFloat(consolidated[key][volumes], 64)
			if err != nil {
				continue
			}
			consolidated[key][volumes] = strconv.FormatFloat(existingValue+value, 'f', -1, 64)
		}
	}

	var result []map[string]string
	for _, record := range consolidated {
		result = append(result, record)
	}
	Data = result
	return result
}
