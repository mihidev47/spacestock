package util

import "math"

func ParseLocationId(locationId string) (location []string) {
	strLen := len(locationId)
	if strLen < 3 {
		return nil
	}
	// Start index
	index := 0
	step := 0
	// Init formatter for Indonesia address
	format := []int{3, 2, 2, 3, 3}
	// Loop
	for index < strLen && step < 5 {
		// Get digit
		index += format[step]
		// extract substring
		tmp := locationId[0:index]
		// Push to location ids
		location = append(location, tmp)
		step++
	}
	return location
}

func Hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
