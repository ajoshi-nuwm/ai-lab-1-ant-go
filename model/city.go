package model

type City struct {
	Name       string
	neighbours map[*City]NeighbourData
}

// struct for neighbour data
type NeighbourData struct {
	Distance, Ferment float64
}

// getter for neighbours
func (this *City) Neighbours() map[*City]NeighbourData {
	return this.neighbours
}

// add city with distance and ferment values
func (this *City) AddNeighbourData(city *City, distance, ferment float64) {
	if this.neighbours == nil {
		this.neighbours = make(map[*City]NeighbourData)
	}
	this.neighbours[city] = NeighbourData{distance, ferment}
}

// string representation
func (this City) String() string {
	return this.Name
}
