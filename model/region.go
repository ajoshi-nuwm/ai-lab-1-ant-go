/*
Package model represents domain model of ant-algorithm problem
City model represents node of a graph with neighbour data in it
Region model represents array of cities and utility methods to interact with
 */
package model

import (
	"strings"
	"math/rand"
	"log"
)

type Region struct {
	cities []*City
}

func (region *Region) Cities() []*City {
	return region.cities
}

// Add cities to region from string representation
func (region *Region) AddCities(citiesData string, neighboursData []string) {
	citiesDataSplitted := strings.Split(citiesData, " ")
	region.cities = make([]*City, len(citiesDataSplitted))
	for i, cityData := range citiesDataSplitted {
		region.cities[i] = &City{Name: cityData}
	}
	for _, v := range neighboursData {
		s := strings.Split(v, " ")
		city1, city2, distance, ferment := s[0], s[1], s[2], s[3]
		region.SearchCity(city1).AddNeighbourDataString(region.SearchCity(city2), distance, ferment)
	}
}

// Search for city by cityName
func (region *Region) SearchCity(cityName string) *City {
	for _, city := range region.cities {
		if city.Name == cityName {
			return city
		}
	}
	log.Fatalf("cannot find city with name = %v", cityName)
	return nil
}

// Get random city to start with
func (region *Region) GetRandomCity() *City {
	return region.cities[rand.Intn(len(region.cities))]
}

// String representation
func (region Region) String() string {
	citiesNames := make([]string, len(region.cities))
	for i, city := range region.cities {
		citiesNames[i] = city.String()
	}
	return strings.Join(citiesNames, "\n\n")
}
