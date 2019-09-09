package main

import (
	"fmt"
)

func main()  {
	array := [6]int{10, 20, 30, 40, 50, 60}
	result := sum(array)
	fmt.Println(result)
}

func sum(array [6]int) int {
	sum := 0
	index := 0

	SumI:
		sum += array[index]
		index++
		if index < len(array) {
			goto SumI
		}
	return sum
}