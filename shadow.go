package main

import (
	"fmt"
)

func main() {

	shadow := "world"
	fmt.Println(shadow)

	{
		// redefine variable in block , will ahve scope only in block
		shadow := "block"
		fmt.Println(shadow)
	}
	fmt.Println("shadow scope not chnaged outside of the block")
	fmt.Println("shadow scope should be avoided")

	fmt.Println(shadow)

}
