package main

import "fmt"

func main() {
  	var (
		r float64
		pi float64 = 3.14159
	)
	
	fmt.Scan(&r)
   	area := pi * (r*r)
	fmt.Printf("A=%.4f\n", area)
}
