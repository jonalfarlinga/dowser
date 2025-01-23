package data

type Flow struct {
	Source      string
	Target      string
	Value       float64
	Topleft     Point
	Bottomleft  Point
	Topright    Point
	Bottomright Point
	Control1    Point
	Control2    Point
}

const (
	TOPL = iota
	TOPR
	BOTL
	BOTR
	CONTROL1
	CONTROL2
)

type Point struct {
	X int
	Y int
}

func (f *Flow) SetPoint(p, x, y int) {
	switch p {
	case TOPL:
		f.Topleft = Point{x, y}
	case TOPR:
		f.Topright = Point{x, y}
	case BOTL:
		f.Bottomleft = Point{x, y}
	case BOTR:
		f.Bottomright = Point{x, y}
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
