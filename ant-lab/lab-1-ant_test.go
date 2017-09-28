package main

import (
	"testing"
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/worker"
)

const (
	antNumber int = 1000
)

func TestAnt(t *testing.T) {
	cases := []struct {
		in   []string
		want float64
	}{{
		[]string{
			"1 2 3 4 5",
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
		},
		171,
	}}

	for _, c := range cases {
		region := model.Region{}
		region.AddCities(c.in[0], c.in[1:])

		visitedCities := []string{}
		var passedDistance float64
		for i := 0; i < antNumber; i++ {
			ant := worker.NewAnt(region.GetRandomCity())
			for ant.VisitCity() {
			}
			if passedDistance == 0 || ant.GetPassedDistance() < passedDistance {
				visitedCities = ant.GetVisitedCities()
				passedDistance = ant.GetPassedDistance()
			}
		}
		if len(visitedCities) != len(region.Cities()) {
			t.Errorf("Ant has not visited all cities. Actual: %v, expected: %v", len(visitedCities), len(region.Cities()))
		}
		if passedDistance != c.want {
			t.Errorf("Ant has not passed through optimal path. Actual: %v, expected: %v", passedDistance, c.want)
		}
	}
}
