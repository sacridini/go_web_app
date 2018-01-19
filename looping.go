package main

import "fmt"

func main() {
	// default
/* 	for i:=0; i < 10; i++ {
		fmt.Println(i)
	} */
	// or

/* 	i:=0
	for i<10 {
		fmt.Println(i)
		i++
	} */

	// infinite loop
/* 	for {
		fmt.Println("dudu")
	} */

	// Simulating a while loop
	x := 5
	for {
		fmt.Println("Do Stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}
}