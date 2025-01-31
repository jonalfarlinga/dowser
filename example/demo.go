package main

import (
	"log"
	"os"
	"water-tracker/data"
	"water-tracker/draw"
)

func main() {
	// log file setup
	file, err := os.OpenFile("output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// get data
	nodeCols := []string{"Source", "Use"}
	volumes := "gals"
	filepath := "10_24_chart.csv"
	err = data.LoadData(filepath)
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}
	data.ConsolidateRecords(nodeCols, volumes)
	nodes, err := data.GetNodes(nodeCols, volumes)
	if err != nil {
		log.Fatalf("Error getting nodes: %v", err)
	}
	flows, err := data.GetFlows(nodeCols, volumes)
	if err != nil {
		log.Fatalf("Error getting flows: %v", err)
	}

	// set nodes and flows positions
	data.SortNodes(nodes)
	data.SetNodesPositions(nodes)
	err = draw.SetFlowsPositions(flows, nodes)
	if err != nil {
		log.Fatalf("Error setting flows positions: %v", err)
	}

	// draw chart
	output := draw.DrawChart(flows, nodes)
	log.Println(output)

	// output to file
	outputBytes := []byte(output)
	err = os.WriteFile("output.svg", outputBytes, 0666)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
}
