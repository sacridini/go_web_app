package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address
	fmt.Println("x =",x)
	fmt.Println("a =",a)
	fmt.Println("*a (Ponteiro) =",*a)

	*a = 5
	fmt.Println("x depois do \"*a = 5\" =",x)

	// complicações e devaneios hahaha
	*a = *a**a 
	fmt.Println("x depois do \"*a = *a**a\" =",x)
	fmt.Println("*a depois do \"*a = *a**a\" =",*a)
}