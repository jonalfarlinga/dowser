package handlers

import (
	"log"
	"net/http"
)

func GetGraph(mux *http.ServeMux) {
	mux.HandleFunc("/api/graph", func(w http.ResponseWriter, r *http.Request) {
		// Call makeGraph to generate the SVG
		svg := makeGraph()

		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(svg))
	})
}

// Dummy implementation of makeGraph
func makeGraph() string {
	// Replace this with the actual implementation
	log.Println("Generating SVG graph...")
	return `<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100">
		<circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" />
	</svg>`
}
