package main

import "fmt"

var global string = "bhaskar"


func main() {
	fmt.Println(global)
	var global string = "modified."
	//avlue declared at local scope level.
	//global := "bhaskarModified"//cannt redclare at the same scope.
	global = "bhaskarModified1"
	fmt.Println(global)
}