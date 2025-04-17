package handlers

import (
	"dowser"
	"encoding/json"
	"net/http"
)

func GetGraph(mux *http.ServeMux) {
	mux.HandleFunc("/api/graph", func(w http.ResponseWriter, r *http.Request) {
		// get input data from request
		var requestData struct {
			NodeCols []string `json:"node_cols"`
			Volumes  string   `json:"volumes"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Call makeGraph to generate the SVG
		svg := dowser.MakeGraph(requestData.NodeCols, requestData.Volumes)

		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(svg))
	})
}

// curl -X POST http://localhost:8080/api/graph `
// -H "Content-Type: application/json" `
// -d '{
//   "node_cols": ["Source", "Use"],
//   "volumes": "gals"
// }'
