package model

import (
	"strconv"
	"fmt"
	"strings"
)

type City struct {
	Name       string
	neighbours map[*City]*NeighbourData
}

// struct for neighbour data
type NeighbourData struct {
	Distance, Ferment float64
}

// getter for neighbours
func (city *City) Neighbours() map[*City]*NeighbourData {
	return city.neighbours
}

// add city with distance and ferment values
func (city *City) AddNeighbourData(neighbour *City, distance, ferment float64) {
	if city.neighbours == nil {
		city.neighbours = make(map[*City]*NeighbourData)
	}
	if neighbour.neighbours == nil {
		neighbour.neighbours = make(map[*City]*NeighbourData)
	}
	city.neighbours[neighbour] = &NeighbourData{distance, ferment}
	neighbour.neighbours[city] = &NeighbourData{distance, ferment}
}

func (city *City) AddNeighbourDataString(neighbour *City, distance, ferment string) {
	d, parseD := strconv.ParseFloat(distance, 64)
	f, parseF := strconv.ParseFloat(ferment, 64)
	if parseD == nil && parseF == nil {
		city.AddNeighbourData(neighbour, d, f)
	} else {
		panic("Can not parse float values")
	}
}

// string representation
func (city City) String() string {
	neighbourData := []string{}
	for city, data := range city.neighbours {
		neighbourData = append(neighbourData, fmt.Sprintf("%v - %v", city.Name, data.String()))
	}
	return fmt.Sprintf("City: %v\nNeighbours:\n%v", city.Name, strings.Join(neighbourData, ",\n"))
}

func (neighbourData NeighbourData) String() string {
	return fmt.Sprintf("distance = %v, ferment = %v", neighbourData.Distance, neighbourData.Ferment)
}
