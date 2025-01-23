package data

type Node struct {
	Label string
	TotalFlow float64
	Level int
	x int
	y int
	width int
	height int
}

type Flow struct {
	Source string
	Target string
	Value float64
}

type CategoryLevel struct {
	Nodes []Node
	TotalVolumes float64
	Level int
}

func NewCategoryLevel(level int) *CategoryLevel {
	return &CategoryLevel{
		Nodes: make([]Node, 0),
		TotalVolumes: 0,
		Level: level,
	}
}

func (cl *CategoryLevel) AddNode(node Node) {
	cl.Nodes = append(cl.Nodes, node)
	cl.TotalVolumes += node.TotalFlow
}
