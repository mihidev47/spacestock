// Package util provides reusable utility functions that is used across packages
package util

import (
	"math/rand"
	"strconv"
	"strings"
)

// Init initialize utility functions. App will exit if an error is occurred while initializing utilities
func Init() {
	// Init error codes
	initErrorCodes()
	// Init snowflake
	initSnowflake()
}

func RandBase36(digit int) (n string) {
	// If digit is less than 0, return
	if digit < 1 {
		return
	}
	// Generate maximum possible number from base36
	maxStr := "Z"
	for i := 0; i < digit-1; i++ {
		maxStr += "Z"
	}
	// Convert to decimal
	max, err := strconv.ParseInt(maxStr, 36, 64)
	if err != nil {
		return
	}
	// Increment max
	max++
	// Generate random number
	var min int64 = 1
	r := min + rand.Int63n(max-min)
	// Convert to base36
	n = strings.ToUpper(strconv.FormatInt(r, 36))
	// Add zerofill
	return Zerofill(n, digit)
}

func Zerofill(n string, digit int) string {
	// If n length is less than digit, add zerofill
	if l := len(n); l < digit {
		diff := digit - l
		for i := 0; i < diff; i++ {
			n = "0" + n
		}
	}
	return n
}
