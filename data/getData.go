package data

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
	"water-tracker/settings"
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

func createNodes(grouped map[string]float64, level int) []*Node {
	var nodes []*Node
	for label, totalFlow := range grouped {
		node := NewNode(label, totalFlow, level)
		nodes = append(nodes, node)
	}
	return nodes
}

func GetNodes(nodeCols []string, volumes string) ([]*Node, error) {
	var nodes []*Node
	for i := range nodeCols {
		grouped := groupByColumn(Data, nodeCols[i], volumes)
		for label, totalFlow := range grouped {
			node := NewNode(label, totalFlow, i)
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
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

func Sort(nodes []*Node) {
	for i := range nodes {
		for j := range nodes {
			if nodes[i].TotalFlow > nodes[j].TotalFlow {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
}

func setLevelNodesPositions(nodes []*Node, levelX int, levelYscalar float64) {
	lastY := 0
	for i := range nodes {
		y := lastY
		nodes[i].SetPosition(levelX, y)
		height := int(nodes[i].TotalFlow * levelYscalar)
		nodes[i].SetHeight(height)
		lastY += height
	}
}

func SetNodesPositions(nodes []*Node) {
	levels := nodeLevels(nodes)
	levelIntervalX := (settings.CHART_WIDTH / (levels - 1)) - settings.NODE_WIDTH
	totalVolumes := make(map[int]float64)
	catLevels := make(map[int][]*Node, levels)
	for i := range nodes {
		if catLevels[i] == nil {
			catLevels[i] = make([]*Node, 0)
		}
		totalVolumes[nodes[i].Level] += nodes[i].TotalFlow
		catLevels[nodes[i].Level] = append(catLevels[nodes[i].Level], nodes[i])
	}
	for i := range catLevels {
		levelX := levelIntervalX * i
		levelYscalar := float64(settings.CHART_HEIGHT) / totalVolumes[nodes[i].Level]
		setLevelNodesPositions(catLevels[i], levelX, levelYscalar)
	}
}

func nodeLevels(nodes []*Node) int {
	levels := make(map[int]bool)
	for i := range nodes {
		levels[nodes[i].Level] = true
	}
	return len(levels)
}
