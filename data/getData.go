package data

import (
	"encoding/csv"
	"fmt"
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

func groupByColumn(records []map[string]string, columnName string) map[string]float64 {
	grouped := make(map[string]float64)
	for _, record := range records {
		key := record[columnName]
		value, err := strconv.ParseFloat(record["gals"], 64)
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

func GetLevels() ([]CategoryLevel, error) {
	levels := make([]CategoryLevel, 0)

	grouped := groupByColumn(Data, "Source")
	nodes := createNodes(grouped, 0)
	levels = append(levels, *NewCategoryLevel(0))
	for _, node := range nodes {
		levels[len(levels)-1].AddNode(node)
		fmt.Printf("Node: %+v\n", node)
	}

	grouped = groupByColumn(Data, "Use")
	nodes = createNodes(grouped, 1)
	levels = append(levels, *NewCategoryLevel(1))
	for _, node := range nodes {
		levels[len(levels)-1].AddNode(node)
		fmt.Printf("Node: %+v\n", node)
	}

	return levels, nil
}

func GetFlows() ([]Flow, error) {
	var flows []Flow
	for _, record := range Data {
		source := record["Source"]
		target := record["Use"]
		value, err := strconv.ParseFloat(record["gals"], 64)
		if err != nil {
			continue
		}
		flow := Flow{
			Source: source,
			Target: target,
			Value:  value,
		}
		flows = append(flows, flow)
		fmt.Printf("Flow: %+v\n", flow)
	}
	return flows, nil
}
