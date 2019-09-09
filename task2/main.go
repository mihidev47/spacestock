package main

import (
	"fmt"
	"strings"
)

func main()  {
	array := [6]string{"agus", "ayam", "tikus", "kodok", "tono", "tiny"}
	search := "ay"
	result, _ := contains(array, search)
	fmt.Println(result)
}

func contains(array [6]string, search string) (string, bool) {
	index := 0

	Search:
		result := strings.Contains(array[index], search)
		if index >= len(array)-1 {
			return "Not Found.", false
		}

		if result == false {
			index++
			goto Search
		}
		
	return array[index], result
}