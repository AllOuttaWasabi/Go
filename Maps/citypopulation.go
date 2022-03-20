package main

import (
	"fmt"
)

func main() {
	/*
		This function main maps out a population to a city and prints out the map.
	*/

	cityPopulation := make(map[string]int)

	cityPopulation["New York"] = 8177025
	cityPopulation["Atlanta"] = 532695
	cityPopulation["Miami"] = 483395
	cityPopulation["Raleigh"] = 474069
	cityPopulation["Richmond"] = 234081

	for index, element := range cityPopulation {
		fmt.Println("City:", index, "\t", "Population:", element)
	}

}
