package main

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/worker"
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/util"
	"fmt"
	"os"
	"strconv"
)

const (
	manFlag = "man"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == manFlag {
		man()
		return
	}

	filename := "in.txt"
	antNumber := 10000
	if len(os.Args) > 2 {
		filename = os.Args[1]
		antNumber, _ = strconv.Atoi(os.Args[2])
	} else if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	data, err := util.ReadFromFile(filename)
	if err != nil {
		man()
		return
	}

	region := model.Region{}
	region.AddCities(data[0], data[1:])

	visitedCities := []string{}
	var passedDistance float64
	var luckyAnt int
	for i := 0; i < antNumber; i++ {
		ant := worker.NewAnt(region.GetRandomCity())
		for ant.VisitCity() {
		}
		if passedDistance == 0 || ant.GetPassedDistance() < passedDistance {
			visitedCities = ant.GetVisitedCities()
			passedDistance = ant.GetPassedDistance()
			luckyAnt = i
		}
	}
	fmt.Printf("visited cities = %v\n", visitedCities)
	fmt.Printf("distance = %v\n", passedDistance)
	fmt.Printf("lucky ant = %v\n", luckyAnt+1)
}

func man() {
	fmt.Println("usage: ant-lab.exe <file-to-analyse | in.txt> [ant-number]")
}
