package main

import "fmt"

func main() {
	var (
		a int
		b int
		total int
	)
	
	fmt.Scanln(&a)
    	fmt.Scanln(&b)
	total = a+b
	fmt.Printf("X = %d\n", total)
}
