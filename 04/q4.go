package main

import (
	"fmt"
	"strings"
)

/*
Exercise 004

Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.
Then, the output should be: [34 67 55 33 12 98]
*/

func main() {
	var input string = "34, 67, 55, 33, 12, 98"
	split := strings.Split(input, ", ")
	fmt.Println(split)
}
