package main

import "fmt"

func main() {

	//Boolean 
	//Numeric - Int, Float, Complex Numbers
	//TextType -- Strings.

	var isThreadRunning bool = false

	fmt.Printf("%v, %T",isThreadRunning,isThreadRunning)


	//
	n := 1 == 1
	m := 1 == 2
	fmt.Println()
	fmt.Printf("%v, %T",n,n)
	fmt.Printf("%v, %T",m,m)



	//Integer Types
	// int8 -128 to 127
	// int16 - -32,768 to 32,767
	// int32
	// int64

	var int8var int8 = 126
	fmt.Printf("%v, %T",int8var,int8var)

	//Unsigned Integers
	// uint8 - 0 - 255
	// uint16 - 

	//Arithmetic Operations
	a:=10
	b:=30

	fmt.Printf("%v\n", a+b)
	fmt.Printf("%v\n", a-b)
	fmt.Printf("%v\n", a*b)
	fmt.Printf("%v\n", a/b)
	fmt.Printf("%v\n", a%b)



	//Type casted...
	var intVar int = 10
	var int8Var int8 = 40

	fmt.Printf("%v\n", intVar+int(int8Var))

	// And Ex Or , and Not


	x:=10
	y:=20
	fmt.Println("--Conditioning Operators --")
	fmt.Printf("%v\n", x&y)
	fmt.Printf("%v\n", x|y)
	fmt.Printf("%v\n", x^y)
	fmt.Printf("%v\n", x&^y)


	fmt.Println("--Bit Shifiting----")


	//k:=8

	//k<<3 //2^3 * 
	//k>>3

	//fmt.Printf("%v\n", k) Not Working

	// FloatingType

	fmt.Println("--FloatingPint Calcuations----")
	d:=12.4
	e:=15.5
	fmt.Println(d+e)

	//1.11 ComplexTypes
	

}