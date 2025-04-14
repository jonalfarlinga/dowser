package draw

import (
	"dowser/data"
	"dowser/settings"
	"fmt"
)

func getNodesPaths(nodes []*data.Node) []string {
	n := make([]string, 0)

	for i := range nodes {
		n = append(
			n, fmt.Sprintf(
				"M %d,%d L %d,%d L %d,%d L %d,%d Z",
				nodes[i].GetPosition().X,
				nodes[i].GetPosition().Y,
				nodes[i].GetPosition().X,
				nodes[i].GetPosition().Y+nodes[i].GetHeight(),
				nodes[i].GetPosition().X+settings.NODE_WIDTH,
				nodes[i].GetPosition().Y+nodes[i].GetHeight(),
				nodes[i].GetPosition().X+settings.NODE_WIDTH,
				nodes[i].GetPosition().Y,
			),
		)
	}
	return n
}

func SetNodesPositions(nodes []*data.Node) {
	levels := numberOfLevels(nodes)
	levelIntervalX := (settings.CHART_WIDTH / (levels - 1)) - settings.NODE_WIDTH
	totalVolumes := make(map[int]float64)
	catLevels := make(map[int][]*data.Node, levels)
	for i := range nodes {
		if catLevels[i] == nil {
			catLevels[i] = make([]*data.Node, 0)
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

func setLevelNodesPositions(nodes []*data.Node, levelX int, levelYscalar float64) {
	lastY := 0
	for i := range nodes {
		y := lastY
		nodes[i].SetPosition(levelX, y)
		height := int(nodes[i].TotalFlow * levelYscalar)
		nodes[i].SetHeight(height)
		lastY += height
	}
}

func numberOfLevels(nodes []*data.Node) int {
	levels := make(map[int]bool)
	for i := range nodes {
		levels[nodes[i].Level] = true
	}
	return len(levels)
}
