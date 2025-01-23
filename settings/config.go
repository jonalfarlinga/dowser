package settings

var (
	CHART_WIDTH  = 1000
	CHART_HEIGHT = 800
	NODE_WIDTH   = 50
)

const (
	DEFAULT_CHART_WIDTH  = 1000
	DEFAULT_CHART_HEIGHT = 800
	DEFAULT_NODE_WIDTH   = 50
)

var (
    FLOWS_COLORS = []string{
        "red",
        "orange",
        "yellow",
        "green",
        "cyan",
        "blue",
        "indigo",
        "violet",
        "purple",
        "magenta",
        "pink",
        "brown",
        "gray",
        "black",
        "white",
    }
    NODES_COLORS = []string{
        "darkred",
        "darkorange",
        "darkyellow",
        "darkgreen",
        "darkcyan",
        "darkblue",
        "darkindigo",
        "darkviolet",
        "darkpurple",
        "darkmagenta",
        "darkpink",
        "darkbrown",
        "darkgray",
        "darkblack",
        "darkwhite",
    }
)

func GetNodesColors(i int) string {
	for i >= len(NODES_COLORS) {
		i -= len(NODES_COLORS)
	}
	return NODES_COLORS[i]
}

func GetFlowsColors(i int) string {
	for i >= len(FLOWS_COLORS) {
		i -= len(FLOWS_COLORS)
	}
	return FLOWS_COLORS[i]
}
