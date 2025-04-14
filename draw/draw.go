package draw

import (
	"dowser/data"
	"dowser/settings"
	"fmt"
	"strconv"
	"strings"
)

func DrawChart(flows []data.Flow, nodes []*data.Node) string {
	components := ""
	flowPaths := getFlowsPaths(flows)
	for i := range flowPaths {
		components += fmt.Sprintf(
			"<path d=\"%s\" fill=\"%s\" fill-opacity=\"0.25\" stroke-width=\"2\" />\n",
			flowPaths[i], getFlowColor(flowPaths[i]),
		)
	}

	nodesPaths := getNodesPaths(nodes)
	for i := range nodesPaths {
		components += fmt.Sprintf(
			"<path d=\"%s\" fill=\"%s\" fill-opacity=\"1.0\" />\n",
			nodesPaths[i], getNodeColor(nodesPaths[i]),
		)
	}

	components += getNodesLabels(nodes)

	return fmt.Sprintf(
		"<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n%s</svg>",
		settings.CHART_WIDTH, settings.CHART_HEIGHT, components,
	)
}

func getFlowColor(path string) string {
	parts := strings.Split(path, " ")
	if len(parts) < 2 {
		return "#000000"
	}
	coords := strings.Split(parts[1], ",")
	if len(coords) < 2 {
		return "#000000"
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return "#000000"
	}
	factor := y * 100 / settings.CHART_HEIGHT
	return fmt.Sprintf(
		"#%02x%02x%02x",
		factor*3/4,
		factor,
		80)
}

func getNodeColor(path string) string {
	parts := strings.Split(path, " ")
	if len(parts) < 2 {
		return "#000000"
	}
	coords := strings.Split(parts[1], ",")
	if len(coords) < 2 {
		return "#000000"
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return "#000000"
	}
	factor := y * 100 / settings.CHART_HEIGHT
	return fmt.Sprintf(
		"#%02x%02x%02x",
		factor,
		factor*2,
		250)
}
