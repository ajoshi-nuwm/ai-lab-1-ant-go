package model

import (
	"strings"
	"fmt"
)

type Region struct {
	cities []*City
}

func (region *Region) Cities() []*City {
	return region.cities
}

func (region *Region) AddCities(citiesData string) {
	citiesDataSplited := strings.Split(citiesData, " ")
	region.cities = make([]*City, len(citiesDataSplited))
	for i, cityData := range citiesDataSplited {
		region.cities[i] = &City{Name: cityData}
	}
}

func (this Region) String() string {
	citiesNames := make([]string, len(this.cities))
	for i, city := range this.cities {
		citiesNames[i] = (*city).Name
	}
	return fmt.Sprintf("{%v}", strings.Join(citiesNames, ", "))
}
