package main

import (
	"fmt"
)

// structure can have methods

type point struct {
	x, y int
}

func main() {
	// Note initialized in curly braces
	p := point{1, 3}
	fmt.Printf("struct is %v\n", p)
	// use %+v to include stuct field names
	fmt.Printf("struct is %+v\n", p)
	// use %#v to print go syntax represnation
	fmt.Printf("struct is %#v\n", p)
	// use type of which object
	fmt.Printf("struct is %T\n", p)
	// pointer represent by &
	fmt.Printf("struct is %p\n", &p)

}
