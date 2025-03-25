package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello()) // It prints "Hello, world."

	// Data Types and Variables
	var (
		decision bool
		semester int
	)

	decision = false
	var gpa float32 = 3.68 // It same as `gpa := 3.68` or `var gpa = 3.68`, etc.
	var name string = "Rafi Abiyyu Airlangga"
	semester = 4

	fmt.Println("The length of my full name is", len(name))
	fmt.Println("The third letter of my name is " + string(name[2]) + ".")

	const grade = 'A'
	fmt.Println("My exam grade is " + string(grade) + ".")

	// Conditional
	if decision {
		fmt.Println("The decision is true.")
		fmt.Println("My name is " + name, "and I'm in semester", semester, "with a GPA of", gpa)
	} else {
		fmt.Println("The decision is false, so I won't print anything.")
	}

	// Looping
	count, sum := 1, 1
	for i := range 10 {
		sum += i
		count++
	}
	fmt.Println("The sum is", sum, "and the count is " + fmt.Sprint(count) + ".")

	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("The sum is now", fmt.Sprint(sum) + ".")

	// Functions
	fmt.Println("Square root of 2025 is", sqrt(2025), "and square root of -4 is " + fmt.Sprint(sqrt(-4)) + ".")
}