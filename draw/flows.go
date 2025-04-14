package draw

import (
	"dowser/data"
	"dowser/settings"
	"fmt"
)

func SetFlowsPositions(flows []data.Flow, nodes []*data.Node) error {
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
		leftHeight := flows[i].Bottomleft.Y - flows[i].Topleft.Y
		rightHeight := flows[i].Bottomright.Y - flows[i].Topright.Y
		if leftHeight < 2 && rightHeight < 2 {
			flows[i].Topleft.Y -= 1
			flows[i].Topright.Y -= 1
			flows[i].Bottomleft.Y += 1
			flows[i].Bottomright.Y += 1 - rightHeight
		}
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
