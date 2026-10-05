package main

import "fmt"

func main() {

	// 1
	myHobbies := [3]string{
		"Gaming", "Coding", "Reading",
	}
	fmt.Println(myHobbies)

	// 2
	fmt.Println(myHobbies[:1])
	fmt.Println(myHobbies[1:])

	// 3
	sliceHobbies := myHobbies[:2]
	fmt.Println(sliceHobbies)

	// 4
	fmt.Println(cap(sliceHobbies))
	sliceHobbies = sliceHobbies[1:3]
	fmt.Println(sliceHobbies)

	// 5
	myGoals := []string{
		"Rich", "Best",
	}
	fmt.Println(myGoals)

	// 6
	myGoals[1] = "Handsome"
	myGoals = append(myGoals, "Nice")
	fmt.Println(myGoals)

	// 7
	type Product struct {
		title string
		id    string
		price float64
	}

	products := []Product{{"t1", "id1", 1.1}, {"t2", "id2", 2.2}}
	fmt.Println(products)

	newProduct := Product{
		"t3", "id3", 3.3,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
