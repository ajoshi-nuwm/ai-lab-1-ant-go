package main

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"fmt"
)

func main() {
	region := model.Region{}
	citiesData := "1 2 3"
	neighboursData := []string {
		"1 2 45 1",
		"1 3 52 1",
		"2 3 32 1",
	}
	region.AddCities(citiesData, neighboursData)
	fmt.Println(region)
}
