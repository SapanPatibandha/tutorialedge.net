package main

import "fmt"

// Flight - a struct that
// contains information about flights
type Flight struct {
  Origin string
  Destination string
  Price int
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
  // implement
  return nil
}

func printFlights(flights []Flight) {
  for _, flight := range flights {
    fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
  }
}

func main() {
  // an empty slice of flights
  var flights []Flight

  sortedList := SortByPrice(flights)
  printFlights(sortedList)
}
