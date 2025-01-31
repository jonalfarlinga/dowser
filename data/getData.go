package data

import (
	"encoding/csv"
	"os"
	"sort"
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

func groupByColumn(records []map[string]string, columnName string, volumes string) map[string]float64 {
	grouped := make(map[string]float64)
	for _, record := range records {
		key := record[columnName]
		value, err := strconv.ParseFloat(record[volumes], 64)
		if err != nil {
			continue
		}
		grouped[key] += value
	}
	return grouped
}

func createNodes(grouped map[string]float64, level int) []Node {
	var nodes []Node
	for label, totalFlow := range grouped {
		node := NewNode(label, totalFlow, level)
		nodes = append(nodes, *node)
	}
	return nodes
}

func GetLevels(nodeCols []string, volumes string) ([]CategoryLevel, error) {
	levels := make([]CategoryLevel, 0)
	for i := range nodeCols {
		grouped := groupByColumn(Data, nodeCols[i], volumes)
		nodes := createNodes(grouped, i)
		levels = append(levels, *NewCategoryLevel(i))
		for _, node := range nodes {
			levels[i].AddNode(node)
		}
	}
	return levels, nil
}

func GetFlows(nodeCols []string, volumes string) ([]Flow, error) {
	var flows []Flow
	for i := 0; i < len(nodeCols)-1; i++ {
		for _, record := range Data {
			source := record[nodeCols[i]]
			target := record[nodeCols[i+1]]
			value, err := strconv.ParseFloat(record[volumes], 64)
			if err != nil {
				continue
			}
			flow := Flow{
				Source: source,
				Target: target,
				Value:  value,
			}
			flows = append(flows, flow)
		}
	}
	sort.Slice(flows, func(i, j int) bool {
		return flows[i].Value > flows[j].Value
	})
	return flows, nil
}
