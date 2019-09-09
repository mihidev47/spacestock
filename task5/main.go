package main

import (
	"fmt"
)

// Case Struct
type Case struct {
	Order int
	Name  string
}

var cases = []Case{
	{Order: 1, Name: "one"},
	{Order: 2, Name: "two"},
	{Order: 3, Name: "tri"}, 
	{Order: 4, Name: "fou"}, 
	{Order: 5, Name: "fiv"},
}

func main()  {
	rearrange(cases, 1, 2)
	rearrange(cases, 3, 1)
}

// rearrange
func rearrange(cases []Case, i int, j int) { 
	idxI, FindI := findOrder(i, cases)
	idxJ, findJ := findOrder(j, cases)

	if FindI == true && findJ == true {
		cases[idxI].Name, cases[idxJ].Name = cases[idxJ].Name, cases[idxI].Name 
	}

	fmt.Println(cases)
}

func findOrder(order int, cases []Case) (int, bool) {
	for i, v := range cases {
		if v.Order == order {
			return i, true
		}
	}
	return 0, false
}
