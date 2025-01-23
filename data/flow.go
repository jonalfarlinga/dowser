package data

type Flow struct {
	Source      string
	Target      string
	Value       float64
	Topleft     Point
	Bottomleft  Point
	Topright    Point
	Bottomright Point
}

const (
	TOPL = iota
	TOPR
	BOTL
	BOTR
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
