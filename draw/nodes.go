package draw

import (
	"fmt"
	"log"
	"sort"
	"water-tracker/data"
	"water-tracker/settings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type Label struct {
	X, Y         int
	nodeX, nodeY int
	Text         string
}

func getNodesPaths(nodes []data.Node) []string {
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

func getNodesLabels(nodes []data.Node) string {
	labels := make([]Label, 0)
	for i := range nodes {
		midY := nodes[i].GetPosition().Y + nodes[i].GetHeight()/2
		edgeX := nodes[i].GetPosition().X + settings.NODE_WIDTH
		textX := edgeX
		if edgeX >= settings.CHART_WIDTH {
			edgeX = nodes[i].GetPosition().X
			textX = edgeX - settings.NODE_WIDTH
		}
		labels = append(
			labels,
			Label{
				X:     textX,
				Y:     midY,
				nodeX: edgeX,
				nodeY: midY,
				Text:  nodes[i].Label,
			})
	}

	labels = checkTextOverlap(labels)
	log.Printf("labels: %+v\n", labels)

	text := ""
	lines := ""
	for i := range labels {
		text += fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-size=\"12\">%s</text>\n",
			labels[i].X, labels[i].Y, labels[i].Text,
		)
		if labels[i].nodeY != labels[i].Y {
			lines += fmt.Sprintf(
				"<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" stroke=\"black\" />\n",
				labels[i].X, labels[i].Y, labels[i].nodeX, labels[i].nodeY,
			)
		}
	}
	return text + lines
}

func checkTextOverlap(labels []Label) []Label {
	adjusted := true
	sort.Slice(labels, func(i, j int) bool {
		return labels[i].Y < labels[j].Y
	})
	for adjusted {
		adjusted = false
		nextY := make(map[int]int)
		for i := 0; i < len(labels); i++ {
			if labels[i].Y > settings.CHART_HEIGHT/2 {
				continue
			}
			if labels[i].Y < getY(nextY, labels[i].X) {
				labels[i].Y = getY(nextY, labels[i].X)
				adjusted = true
			}
			putY(nextY, labels[i].X, labels[i].Y+10)
		}
		nextY = endY(nextY)
		for i := len(labels) - 1; i >= 0; i-- {
			if labels[i].Y < settings.CHART_HEIGHT/2 {
				continue
			}
			if labels[i].Y > getY(nextY, labels[i].X) {
				labels[i].Y = getY(nextY, labels[i].X)
				adjusted = true
			}
			putY(nextY, labels[i].X, labels[i].Y-10)
		}
	}
	return labels
}

func putY(m map[int]int, key int, value int) {
	if _, exists := m[key]; !exists {
		m[key] = 0
	}
	m[key] = value
}

func getY(m map[int]int, key int) int {
	if _, exists := m[key]; !exists {
		return -1
	}
	return m[key]
}

func endY(m map[int]int) map[int]int {
	for k := range m {
		m[k] = settings.CHART_HEIGHT
	}
	return m
}

func getTextWidth(text string) int {
	face := basicfont.Face7x13
	d := &font.Drawer{
		Face: face,
	}
	width := d.MeasureString(text)
	return width.Ceil()
}
