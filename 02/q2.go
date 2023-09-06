package main

import (
	"fmt"
	"log"
)

/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

func main() {
	fmt.Print("Enter an integer number: ")

	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Oh no, oopsy")
	}

	result, err := factorial(input)
	if err != nil {
		log.Fatalf("some problem for input %v: %v", input, err)
	}

	fmt.Printf("Factorial of %v is %v", input, result)
}

func factorial(a int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("no factorial for negatives")
	}

	if a == 0 {
		return 1, nil
	}

	fact := 1
	for i := fact; i <= a; i++ {
		fact *= i
	}
	return fact, nil
}
