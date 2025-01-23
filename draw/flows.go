package draw

import (
	"fmt"
	"water-tracker/data"
	"water-tracker/settings"
)

func SetFlowsPositions(flows []data.Flow, nodes []data.Node) error {
	err := fmt.Errorf("Must SetNodePositions before SetFlowsPositions")
	for i := range nodes {
		if nodes[i].GetPos().Y != 0 {
			err = nil
			break
		}
	}
	if err != nil {
		return err
	}

	for _, flow := range flows {
		source := flow.GetSource(nodes)
		target := flow.GetTarget(nodes)
		flow.SetPoint(
			data.START,
			source.GetPos().X+settings.NODE_WIDTH,
			source.GetPos().Y,
		)
		flow.SetPoint(
			data.END,
			target.GetPos().X,
			target.GetPos().Y+target.GetHeight(),
		)
		width := target.GetPos().X - source.GetPos().X
		flow.SetPoint(
			data.CONTROL1,
			(source.GetPos().X+width/3),
			source.GetPos().Y,
		)
		flow.SetPoint(
			data.CONTROL2,
			(target.GetPos().X-width/3),
			source.GetPos().Y,
		)
	}
	return nil
}
