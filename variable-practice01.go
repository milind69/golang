package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var xs int
	fmt.Printf("default inited to %d", xs)

	var e int
	e = 1
	fmt.Println(e)

	// := shorthand, also when doing shorthan no type def needed
	s := "apple"
	fmt.Println(s)

	// bool
	var t bool
	t = false
	fmt.Println(t)

}
