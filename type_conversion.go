package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "33"
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of s is %T  and values is %v\n", s, s)
	fmt.Printf("The type of i is %T  and values is %v\n", i, i)
}
