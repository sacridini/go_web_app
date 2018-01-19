package main

import ("fmt"
		"math/rand")

// Random generation number funtion 
func randomGen() {
	fmt.Println("A number from 1-100:", rand.Intn(100))
}

func add(x float64,y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a,b
}

func main() {	
	// Declaração de variáveis 
	// var num1 float64 = 5.63
	// var num2 float64 = 7.22
	// ou
	// var num3, num4 float64 = 1.26, 99.21
	// ou
	num5, num6 := 6.78, 10.24
	fmt.Println(add(num5,num6))

	// Multiple function
	w1, w2  := "Hey", "There"
	fmt.Println(multiple(w1,w2))
}

