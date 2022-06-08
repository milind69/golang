package main

import (
	"fmt"
)

func main() {
	// inline coditional note ; at the end
	// variable will only be available in block where declared
	if x := 1; x == 1 {
		fmt.Println(x)
	}
}
