package data

import (
	"water-tracker/settings"
)

type Node struct {
	Label     string
	TotalFlow float64
	Level     int
	point     Point
	width     int
	height    int
}

func NewNode(label string, tf float64, level int) *Node {
	return &Node{
		Label:     label,
		TotalFlow: tf,
		Level:     level,
		width:     settings.NODE_WIDTH,
	}
}

func (n *Node) SetPosition(x, y int) {
	n.point = Point{x, y}
}

func (n *Node) GetPos() Point {
	return n.point
}

func (n *Node) SetHeight(h int) {
	n.height = h
}

func (n *Node) GetHeight() int {
	return n.height
}
