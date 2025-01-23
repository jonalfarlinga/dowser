package data

type Flow struct {
	Source   string
	Target   string
	Value    float64
	Start    Point
	End      Point
	Control1 Point
	Control2 Point
}

const (
	START = iota
	END
	CONTROL1
	CONTROL2
)

type Point struct {
	X int
	Y int
}

func (f *Flow) SetPoint(p, x, y int) {
	switch p {
	case START:
		f.Start = Point{x, y}
	case END:
		f.End = Point{x, y}
	case CONTROL1:
		f.Control1 = Point{x, y}
	case CONTROL2:
		f.Control2 = Point{x, y}
	}
}

func (f *Flow) GetSource(nodes []Node) *Node {
	for i := range nodes {
		if nodes[i].Label == f.Source {
			return &nodes[i]
		}
	}
	return nil
}

func (f *Flow) GetTarget(nodes []Node) *Node {
	for i := range nodes {
		if nodes[i].Label == f.Target {
			return &nodes[i]
		}
	}
	return nil
}
