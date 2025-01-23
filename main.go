package main

import (
	"fmt"
	"water-tracker/data"
)

func main() {
	err := data.LoadData("data/test_data.csv")
	if err != nil {
		panic(err)
	}
	catsLevels, err := data.GetLevels()
	if err != nil {
		panic(err)
	}
	flows, err := data.GetFlows()
	if err != nil {
		panic(err)
	}

	for i, level := range catsLevels {
		fmt.Printf("Level %d: %+v\n", i, level)
	}
	fmt.Printf("Flows: %+v\n", flows)
}
