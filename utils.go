package main

import (
	"fmt"
	"math"
)

func sqrt(number float64) string {
	if (number < 0) {
		return sqrt(-number) + "i"
	}
	return fmt.Sprint(math.Sqrt(number))
}