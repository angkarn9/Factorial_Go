package main

import (
	"fmt"

	"./factorial"
)

func main() {
	num := 5
	fmt.Printf("%d! = %d\n", num, factorial.Factorial(num))
}
