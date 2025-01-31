package data

import (
	"sort"
	"strconv"
	"water-tracker/settings"
)

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

func SortNodes(nodes []*Node) {
	for i := range nodes {
		for j := range nodes {
			if nodes[i].TotalFlow > nodes[j].TotalFlow {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
}

func SetNodesPositions(nodes []*Node) {
	levels := numberOfLevels(nodes)
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

func numberOfLevels(nodes []*Node) int {
	levels := make(map[int]bool)
	for i := range nodes {
		levels[nodes[i].Level] = true
	}
	return len(levels)
}
