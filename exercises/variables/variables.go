package main

import "fmt"

func main() {
	// Expliclity defining the variable types for a-d
	var a string = "string"
	fmt.Println("a = ", a)

	var b, c int = 1, 2
	fmt.Println("b and c = ",  b, c)

	var d bool = true
	fmt.Println("d = ", d)

	// Inferring types for e and f
	var e = "inferred string"
	fmt.Println("e = ", e)

	var f, g = 3, 4
	fmt.Println("f and g = ", f,g )

	// Zero-valued - no initialization declared for h and i
	var h int 
	fmt.Println("h = ", h)
	
	var i string 
	fmt.Println("i = ", i)

	// Short hand declared for initializing var
	j := "shorthand string"
	fmt.Println("j = ", j)
}