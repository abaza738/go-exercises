package main

import (
	"fmt"
	"log"
)

/*
Exercise 003

With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8
Then, the output should be:
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

*/

func main() {
	fmt.Print("Enter an integer number: ")

	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Oh no, oopsy")
	}

	result, err := generate_map(input)
	if err != nil {
		log.Fatalf("some problem for input %v: %v", input, err)
	}

	fmt.Println(result)
}

func generate_map(a int) (map[int]int, error) {
	output := make(map[int]int)
	for i := 1; i <= a; i++ {
		output[i] = i * i
	}
	return output, nil
}
