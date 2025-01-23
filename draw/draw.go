package draw

import (
	"fmt"
	"strconv"
	"strings"
	"water-tracker/data"
	"water-tracker/settings"
)

func SetFlowsPositions(flows []data.Flow, nodes []data.Node) error {
	err := fmt.Errorf("Must SetNodePositions before SetFlowsPositions")
	for i := range nodes {
		if nodes[i].GetPosition().Y != 0 {
			err = nil
			break
		}
	}
	if err != nil {
		return err
	}

	sourceY := make(map[string]int)
	targetY := make(map[string]int)

	for i := range flows {
		source := flows[i].GetSource(nodes)
		if _, exists := sourceY[source.Label]; !exists {
			sourceY[source.Label] = 0
		}
		target := flows[i].GetTarget(nodes)
		if _, exists := targetY[target.Label]; !exists {
			targetY[target.Label] = 0
		}

		heightS := float64(source.GetHeight()) / source.TotalFlow * flows[i].Value
		heightT := float64(target.GetHeight()) / target.TotalFlow * flows[i].Value

		flows[i].SetPoint(
			data.TOPL,
			source.GetPosition().X+settings.NODE_WIDTH,
			source.GetPosition().Y+sourceY[source.Label],
		)
		sourceY[source.Label] += int(heightS)
		flows[i].SetPoint(
			data.BOTL,
			source.GetPosition().X+settings.NODE_WIDTH,
			source.GetPosition().Y+sourceY[source.Label],
		)
		flows[i].SetPoint(
			data.TOPR,
			target.GetPosition().X,
			target.GetPosition().Y+targetY[target.Label],
		)
		targetY[target.Label] += int(heightT)
		flows[i].SetPoint(
			data.BOTR,
			target.GetPosition().X,
			target.GetPosition().Y+targetY[target.Label],
		)
	}
	return nil
}

func getFlowsPaths(flows []data.Flow) []string {
	f := make([]string, 0)
	for i := range flows {
		width := flows[i].Topright.X - flows[i].Topleft.X
		control1T := data.Point{
			X: flows[i].Topleft.X + width/2,
			Y: flows[i].Topleft.Y,
		}
		control2T := data.Point{
			X: flows[i].Topright.X - width/2,
			Y: flows[i].Topright.Y,
		}
		control2B := data.Point{
			X: flows[i].Bottomright.X - width/2,
			Y: flows[i].Bottomright.Y,
		}
		control1B := data.Point{
			X: flows[i].Bottomleft.X + width/2,
			Y: flows[i].Bottomleft.Y,
		}
		f = append(
			f, fmt.Sprintf(
				"M %d,%d C %d,%d %d,%d %d,%d L %d,%d C %d,%d %d,%d %d,%d Z",
				flows[i].Topleft.X, flows[i].Topleft.Y,
				control1T.X, control1T.Y,
				control2T.X, control2T.Y,
				flows[i].Topright.X, flows[i].Topright.Y,
				flows[i].Bottomright.X, flows[i].Bottomright.Y,
				control2B.X, control2B.Y,
				control1B.X, control1B.Y,
				flows[i].Bottomleft.X, flows[i].Bottomleft.Y,
			),
		)
	}
	return f
}

func getNodesPaths(nodes []data.Node) []string {
	n := make([]string, 0)
	for i := range nodes {
		n = append(
			n, fmt.Sprintf(
				"M %d,%d L %d,%d L %d,%d L %d,%d Z",
				nodes[i].GetPosition().X, nodes[i].GetPosition().Y,
				nodes[i].GetPosition().X, nodes[i].GetPosition().Y+nodes[i].GetHeight(),
				nodes[i].GetPosition().X+settings.NODE_WIDTH, nodes[i].GetPosition().Y+nodes[i].GetHeight(),
				nodes[i].GetPosition().X+settings.NODE_WIDTH, nodes[i].GetPosition().Y,
			),
		)
	}
	return n
}

func DrawChart(flows []data.Flow, nodes []data.Node) string {
	components := ""
	flowPaths := getFlowsPaths(flows)
	for i := range flowPaths {
		components += fmt.Sprintf(
			"<path d=\"%s\" fill=\"%s\" fill-opacity=\"0.5\" />\n",
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
		factor,
		factor*2,
		180)
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
		120)
}
