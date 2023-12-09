package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type ByPrice []Flight

// func (p ByPrice) Len() int {
// 	return len(p)
// }

func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPrice) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}
func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d \n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	flights := []Flight{
		Flight{Price: 30},
		Flight{Price: 20},
		Flight{Price: 50},
		Flight{Price: 1000},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
