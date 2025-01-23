package data

import "water-tracker/settings"

type CategoryLevel struct {
	Nodes        []Node
	TotalVolumes float64
	Level        int
}

func NewCategoryLevel(level int) *CategoryLevel {
	return &CategoryLevel{
		Nodes:        make([]Node, 0),
		TotalVolumes: 0,
		Level:        level,
	}
}

func (cl *CategoryLevel) AddNode(node Node) {
	cl.Nodes = append(cl.Nodes, node)
	cl.TotalVolumes += node.TotalFlow
}

func (cl *CategoryLevel) SetNodePositions(levels int) {
	levelIntervalX := (settings.CHART_WIDTH / (levels - 1)) - settings.NODE_WIDTH
	levelX := levelIntervalX * cl.Level
	levelScalarY := float64(settings.CHART_HEIGHT) / cl.TotalVolumes
	lastY := 0
	for i := range cl.Nodes {
		y := lastY
		cl.Nodes[i].SetPosition(levelX, y)
		height := int(cl.Nodes[i].TotalFlow * levelScalarY)
		cl.Nodes[i].SetHeight(height)
		lastY += height
	}
}

func (cl *CategoryLevel) Sort() {
	for i := range cl.Nodes {
		for j := range cl.Nodes {
			if cl.Nodes[i].TotalFlow > cl.Nodes[j].TotalFlow {
				cl.Nodes[i], cl.Nodes[j] = cl.Nodes[j], cl.Nodes[i]
			}
		}
	}
}
