package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	// courseRatings := map[string]float64{}

	courseRatings["go"] = 4.6
	courseRatings["react"] = 4.7
	courseRatings["angular"] = 4.8
	courseRatings["node"] = 4.9

	courseRatings.output()

	// fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
