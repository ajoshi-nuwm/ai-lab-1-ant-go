package main

import (
	"github.com/ajoshi-nuwm/ai-lab-1-ant-go/model"
	"fmt"
)

func main() {
	region := model.Region{}
	region.AddCities("1 2 3")
	fmt.Println(region)
}
