package main

import (
	"fmt"
)

func main() {
	var myFloat32 float32 = 4.5
	var myComplex64 complex64 = 4.5
	fmt.Println(myComplex64)
	fmt.Println(myFloat32)

	type Allstring string // create alias for the type string
	const myString = "hello"
	var zString Allstring = myString
	fmt.Println(zString)

	const Aint = 1
	var myFloat = Aint //this will work as it a untyped const
	fmt.Println(myFloat)
	var myflt float64 = Aint
	fmt.Println(myflt) // this will not fail to
	/*
		const typedInt int = 1
		var xFloat64 float64 = typedInt // will fail as cont is typed constant
		fmt.Println(xFloat64)
	*/
}
