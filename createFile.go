package main

import "os"

func main() {
	/*
		Function main takes an array of city names and puts them into a txt file named Eastern_US_Cities.
	*/

	var cities = [5]string{"New York", "Atlanta", "Miami", "Raleigh", "Richmond"}

	file, err := os.Create("Eastern_US_Cities.txt")

	if err != nil {
		return
	}

	defer file.Close()

	for i := 0; i < 5; i++ {
		file.WriteString(cities[i] + "\n")
	}

}
