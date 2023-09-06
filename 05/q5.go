package main

import (
	"fmt"
)

/*
Create a seperate file (module) which has at least two methods:

- ReadString: to read a string from console input
- PrintString: to return the string in upper case.

Also create a `main.go` file that acts as calling class.
*/

func main() {
	ReadString()
	out := PrintString()
	fmt.Println(out)
}
