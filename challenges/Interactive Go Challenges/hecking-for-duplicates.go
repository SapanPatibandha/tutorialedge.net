package main

import "fmt"

// type Developer struct {
// 	Name string
// 	Age  int
// }

func FilterUnique(developers []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueNames []string

	for _, dev := range developers {
		if _, exists := uniqueMap[dev]; !exists {
			uniqueMap[dev] = true
			uniqueNames = append(uniqueNames, dev)
		}
	}
	return uniqueNames
}

func main() {
	fmt.Println("Filter Unique Challenge")

	name := []string{
		"Elliot",
		"Sapan",
		"Alan",
		"Jennifer",
		"Graham",
		"Paul",
		"Alan",
		"Sapan",
	}

	uniqueNames := FilterUnique(name)
	fmt.Println(uniqueNames)
}
