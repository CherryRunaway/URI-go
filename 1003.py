package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	
	fmt.Scan(&a)
  fmt.Scan(&b)
	soma := a+b
	fmt.Printf("SOMA = %d\n", soma)
}
