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

	err = data.LoadData("10_24_chart.csv")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}
    data.ConsolidateRecords(data.Data, []string{"Source", "Use"})
	catsLevels, err := data.GetLevels()
	if err != nil {
		log.Fatalf("Error getting levels: %v", err)
	}
	flows, err := data.GetFlows()
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

	for i := range catsLevels {
        for j := range catsLevels[i].Nodes {
            log.Printf("Node %d: %s\n", j, catsLevels[i].Nodes[j].Label)
        }
	}
    for i := range flows {
        log.Printf("Flow %d: %s-%s-%f\n", i, flows[i].Source, flows[i].Target, flows[i].Value)
    }

	log.Print(draw.DrawChart(flows, nodes))
}
