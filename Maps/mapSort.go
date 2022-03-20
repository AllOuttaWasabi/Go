package main

import (
	"fmt"
	"sort"
)

func main() {

	ChickFilAPerCity := map[string]int{"Raleigh": 20, "Greensboro": 10, "Wilmington": 6, "Winston-Salem": 9, "Greenville": 4}

	fmt.Println(ChickFilAPerCity)

	keys := make([]string, 0, len(ChickFilAPerCity))

	for key := range ChickFilAPerCity {
		keys = append(keys, key)
	}

	// Sort based on key order.
	sort.Strings(keys)
	// Check out the keys sorted with the print statement.
	fmt.Println("Map sorted by key name: ", keys)

	elements := make([]int, 0, len(ChickFilAPerCity))

	for _, element := range ChickFilAPerCity {
		elements = append(elements, element)
	}

	// Sort based on element (value) order.
	sort.Ints(elements)
	// Check out the values with the print statement.
	fmt.Println("Map sorted by element: ", elements)

}
