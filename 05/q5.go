package main

import (
	"fmt"
	"go-exercises/05/g"
)

/*
Create a seperate file (module) which has at least two methods:

- ReadString: to read a string from console input
- PrintString: to return the string in upper case.

Also create a `main.go` file that acts as calling class.
*/

func main() {
	g.ReadString()
	out := g.PrintString()
	fmt.Println(out)
}
