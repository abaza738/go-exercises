package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included).  The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

func main() {
	res, err := solution(2000, 3200)

	if err != nil {
		fmt.Println("Oepsie doepsie!")
		return
	}

	fmt.Println(strings.Join(res, ","))
}

func solution(a int, b int) ([]string, error) {
	var res []string
	for i := a; i < b; i++ {
		if i%5 == 0 && i%7 == 0 {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res, nil
}
