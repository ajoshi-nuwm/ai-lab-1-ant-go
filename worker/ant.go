/*
Package worker represents implementation of ant, which is worker for the algorithm

Ant has constructor to use in most of cases which accepts pointer to city from which ant starts

Utility struct Segment used to represent probability rule to get random city to visit
 */
package worker

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"math/rand"
)

const (
	tripWeight        = 100
	forgetCoefficient = .1
)

type Ant struct {
	visited []*model.City
}

// Segment of probability ruler
type Segment struct {
	city  *model.City
	start float64
	end   float64
}

// Checks if number is in segment
func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}

// Constructor
func NewAnt(initialCity *model.City) *Ant {
	ant := Ant{}
	ant.visited = append(ant.visited, initialCity)
	return &ant
}

// Returns visited cities
func (ant *Ant) GetVisitedCities() []string {
	cities := []string{}
	for _, city := range ant.visited {
		cities = append(cities, city.Name)
	}
	return cities
}

// Visits city and returns true if succeeded, returns false if no available city to visit - end of algorithm
func (ant *Ant) VisitCity() bool {
	visitCity := ant.getVisitCity()
	if visitCity == nil {
		ant.spreadFerment()
		return false
	}
	ant.visited = append(ant.visited, visitCity)
	return true
}

// Get distance passed by ant
func (ant *Ant) GetPassedDistance() float64 {
	var distance float64
	for i := 0; i < len(ant.visited)-1; i++ {
		currentCity := ant.visited[i]
		nextCity := ant.visited[i+1]
		distance += currentCity.Neighbours()[nextCity].Distance
	}
	return distance
}

// Returns next city to visit
func (ant *Ant) getVisitCity() *model.City {
	segments := []Segment{}
	var start float64
	for _, city := range ant.getAvailableCities() {
		probability := ant.calculateCityTransitionProbability(city)
		segments = append(segments, Segment{city, start, start + probability})
		start += probability
	}
	random := rand.Float64() * start
	for _, segment := range segments {
		if segment.isInSegment(random) {
			return segment.city
		}
	}
	return nil
}

// Returns all cities from current one to visit except black list - list of visited ones
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

// Returns probability to visit given city from current one
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

// Post-step action - spreading ferment through all visited roads
func (ant *Ant) spreadFerment() {
	distance := ant.GetPassedDistance()
	delta := tripWeight / distance
	for i := 0; i < len(ant.visited)-1; i++ {
		currentCity := ant.visited[i]
		nextCity := ant.visited[i+1]
		oldFerment := (currentCity.Neighbours()[nextCity]).Ferment

		(currentCity.Neighbours()[nextCity]).Ferment = oldFerment*forgetCoefficient + delta
		(nextCity.Neighbours()[currentCity]).Ferment = oldFerment*forgetCoefficient + delta
	}
}

// Checks if city contains in slice
func cityContains(cities [] *model.City, city *model.City) bool {
	for _, c := range cities {
		if c == city {
			return true
		}
	}
	return false
}
