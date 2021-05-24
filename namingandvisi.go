package main

import "fmt"
import "strconv"

//local scope.
var a int = 12

//GlobalScope. Use Capital Letter
var A int = 12

func main() {

	var a int = 12
	fmt.Println(a)

	//Using naming convention.


	//Using Type Convention
	fmt.Printf("%v,%T",a,a)

	var b float32
	b = float32(a)
	fmt.Printf("%v,%T",b,b)


	//String Conversion
	// Here a converted from int to string
	var c string
	c = strconv.Itoa(a)
	fmt.Println()
	fmt.Printf("%v,%T",c,c)
}