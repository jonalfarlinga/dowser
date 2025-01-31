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
	x, y         int
	nodeX, nodeY int
	leftAnchor   bool
	text         string
}

func getNodesLabels(nodes []*data.Node) string {
	labels := make([]*Label, 0)
	for i := range nodes {
		midY := nodes[i].GetPosition().Y + nodes[i].GetHeight()/2
		edgeX := nodes[i].GetPosition().X + settings.NODE_WIDTH
		textX := edgeX
		leftAnchor := true
		if edgeX >= settings.CHART_WIDTH {
			edgeX = nodes[i].GetPosition().X
			textX = edgeX - getTextWidth(nodes[i].Label)
			leftAnchor = false
		}
		labels = append(
			labels,
			&Label{
				x:          textX,
				y:          midY,
				nodeX:      edgeX,
				nodeY:      midY,
				leftAnchor: leftAnchor,
				text:       nodes[i].Label,
			})
	}

	labels = checkTextOverlap(labels)
	log.Printf("labels: %+v\n", labels)

	text := ""
	lines := ""
	for i := range labels {
		lines += connectNodeLabel(labels[i])

		text += fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-size=\"12\">%s</text>\n",
			labels[i].x, labels[i].y, labels[i].text,
		)
		log.Println(labels[i].text, labels[i].x, labels[i].y, labels[i].nodeX, labels[i].nodeY)
	}
	return text + lines
}

func connectNodeLabel(label *Label) string {
	if label.nodeY == label.y {
		return ""
	}
	label.x += settings.NODE_WIDTH
	anchorX := label.x
	if !label.leftAnchor {
		label.x -= settings.NODE_WIDTH + 25
		anchorX = label.x + getTextWidth(label.text) - 5
	}
	labelY := label.y - 5
	if labelY < 15 {
		labelY = 15
	}
	line := fmt.Sprintf(
		"<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" stroke=\"black\" />\n",
		anchorX, labelY, label.nodeX, label.nodeY,
	)
	log.Println(line)

	return line
}

func checkTextOverlap(labels []*Label) []*Label {
	adjusted := true
	sort.Slice(labels, func(i, j int) bool {
		return labels[i].y < labels[j].y
	})
	for adjusted {
		adjusted = false
		nextY := make(map[int]int)
		for i := 0; i < len(labels); i++ {
			if labels[i].y > settings.CHART_HEIGHT/2 {
				continue
			}
			if labels[i].y < getY(nextY, labels[i].nodeX) {
				labels[i].y = getY(nextY, labels[i].nodeX)
				adjusted = true
			}
			putY(nextY, labels[i].nodeX, labels[i].y+10)
		}
		nextY = endY(nextY)
		for i := len(labels) - 1; i >= 0; i-- {
			if labels[i].y < settings.CHART_HEIGHT/2 {
				continue
			}
			if labels[i].y > getY(nextY, labels[i].nodeX) {
				labels[i].y = getY(nextY, labels[i].nodeX)
				adjusted = true
			}
			putY(nextY, labels[i].nodeX, labels[i].y-10)
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
		m[k] = settings.CHART_HEIGHT - 5
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
