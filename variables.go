package main

import "fmt"

//no private scope for variable.
//lower case first letter for package scope
// upper case first letter for export.

// package level variable
// you cant use := operations , it only detects full declaration
var a1 int = 200

var (
	id int = 10
	fname string = "Bhaskar"
	lname string = "s"
)

// 3 level declaration.
func main() {
	//later
	var a int
	a = 100
	fmt.Println(a)

	//direct
	var b int = 500
	fmt.Println(b)

	//direct assing
	c:=4
	fmt.Println(c)

	fmt.Printf("%T, %v",c,c)
	fmt.Printf("%T, %v",c,c)

	fmt.Println()
	fmt.Printf("%T,%v",id,id)
	fmt.Printf("%T,%v",fname,fname)
}