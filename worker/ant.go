package worker

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"math/rand"
	"fmt"
)

const (
	tripWeight        = 100
	forgetCoefficient = .1
)

type Ant struct {
	visited []*model.City
}

// segment of probability ruler
type Segment struct {
	city  *model.City
	start float64
	end   float64
}

func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}

func NewAnt(initialCity *model.City) *Ant {
	ant := Ant{}
	ant.visited = append(ant.visited, initialCity)
	return &ant
}

func (ant *Ant) VisitCity() bool {
	visitCity := ant.getVisitCity()
	if visitCity == nil {
		ant.spreadFerment()
		return false
	}
	ant.visited = append(ant.visited, visitCity)
	return true
}

// get distance passed by ant
func (ant *Ant) GetPassedDistance() float64 {
	var distance float64
	for i := 0; i < len(ant.visited)-1; i++ {
		currentCity := ant.visited[i]
		nextCity := ant.visited[i+1]
		distance += currentCity.Neighbours()[nextCity].Distance
	}
	return distance
}

func (ant *Ant) getVisitCity() *model.City {
	segments := []Segment{}
	var start float64
	for _, city := range ant.getAvailableCities() {
		probability := ant.calculateCityTransitionProbability(city)
		segments = append(segments, Segment{city, start, start + probability})
		start += probability
	}
	random := rand.Float64()
	for _, segment := range segments {
		if segment.isInSegment(random) {
			return segment.city
		}
	}
	fmt.Printf("random = %v", random)
	fmt.Printf("segments = %v", segments)
	return nil
}

func (ant *Ant) getAvailableCities() []*model.City {
	currentCity := ant.visited[len(ant.visited)-1]
	availableCities := []*model.City{}
	for city := range currentCity.Neighbours() {
		if !cityContains(ant.visited, city) {
			availableCities = append(availableCities, city)
		}
	}
	return availableCities
}

func (ant *Ant) calculateCityTransitionProbability(city *model.City) float64 {
	currentCity := ant.visited[len(ant.visited)-1]
	var sum float64
	for _, neighbourData := range currentCity.Neighbours() {
		sum += neighbourData.Ferment / neighbourData.Distance
	}
	candidateData := currentCity.Neighbours()[city]
	candidateValue := candidateData.Ferment / candidateData.Distance
	return candidateValue / sum
}

func (ant *Ant) spreadFerment() {
	distance := ant.GetPassedDistance()
	delta := tripWeight / distance
	for i := 0; i < len(ant.visited)-1; i++ {
		currentCity := ant.visited[i]
		nextCity := ant.visited[i+1]
		oldFerment := (&currentCity.Neighbours()[nextCity]).Ferment

		(&currentCity.Neighbours()[nextCity]).Ferment = oldFerment*forgetCoefficient + delta
		(&nextCity.Neighbours()[currentCity]).Ferment = oldFerment*forgetCoefficient + delta
	}
}

func cityContains(cities [] *model.City, city *model.City) bool {
	for _, c := range cities {
		if c == city {
			return true
		}
	}
	return false
}
