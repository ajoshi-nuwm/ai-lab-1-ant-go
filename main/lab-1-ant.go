package main

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/worker"
	"fmt"
)

func main() {
	region := model.Region{}
	citiesData := "1 2 3 4 5"
	neighboursData := []string{
		"1 2 38 1",
		"1 3 74 1",
		"1 4 59 1",
		"1 5 45 1",
		"2 3 46 1",
		"2 4 61 1",
		"2 5 72 1",
		"3 4 49 1",
		"3 5 85 1",
		"4 5 42 1",
	}
	region.AddCities(citiesData, neighboursData)
	ant := worker.NewAnt(region.GetRandomCity())
	for ant.VisitCity() {

	}
	fmt.Printf("distance = %v\n", ant.GetPassedDistance())
	fmt.Printf("visited cities = %v\n", ant.GetVisitedCities())
}
