package main

import "fmt"

func main() {
	// create a map
	grades := make(map[string]float32)

	// add things
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	// print it
	fmt.Println(grades["Jess"])
	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println("TimsGrade is: ", TimsGrade)

	// remove a item
	delete(grades, "Timmy")
	fmt.Println(grades)

	// interation on maps
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}

