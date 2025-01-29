package main

import (
	"log"
	"os"
	"water-tracker/data"
	"water-tracker/draw"
)

func main() {
	file, err := os.OpenFile("output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	nodeCols := []string{"Source", "Use"}
	volumes := "gals"
	filepath := "10_24_chart.csv"
	err = data.LoadData(filepath)
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}
	data.ConsolidateRecords(data.Data, nodeCols)
	catsLevels, err := data.GetLevels(nodeCols)
	if err != nil {
		log.Fatalf("Error getting levels: %v", err)
	}
	flows, err := data.GetFlows(nodeCols, volumes)
	if err != nil {
		log.Fatalf("Error getting flows: %v", err)
	}
	for _, level := range catsLevels {
		level.Sort()
		level.SetNodePositions(len(catsLevels))
	}

	nodes := make([]data.Node, 0)
	for _, level := range catsLevels {
		nodes = append(nodes, level.Nodes...)
	}
	err = draw.SetFlowsPositions(flows, nodes)
	if err != nil {
		log.Fatalf("Error setting flows positions: %v", err)
	}

	output := draw.DrawChart(flows, nodes)
	log.Println(output)
	outputBytes := []byte(output)
	err = os.WriteFile("output.svg", outputBytes, 0666)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
}
