package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Exercise 6
Write a program that calculates and prints the value according to the given formula:
Q = Square root of [(2 * C * D)/H]

Following are the fixed values of C and H:
    C is 50. H is 30.
    D is the variable whose values should be input to your program in a comma-separated sequence.

Example:
Let us assume the following comma separated input sequence is given to the program: 100,150,180
The output of the program should be: 18,22,24
*/

var C int = 50
var H int = 30

func main() {
	fmt.Println("Enter sequence: ")

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	numbers := strings.Split(input, ",")
	var out []string

	for _, v := range numbers {
		a, _ := strconv.Atoi(v)
		out = append(out, fmt.Sprintf("%d", formula(a)))
	}

	fmt.Println(strings.Join(out, ","))
}

func formula(in int) int {
	return int(math.Sqrt(float64((2 * C * in) / H)))
}
