package main

import "fmt"

func main(){
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int 
	fmt.Println(e)

	f:= "apple"
	fmt.Println(f)
}


/*
var declares 1 or more variables 
Can declare multiple variables 
Go will infer the type of initalized variable 

Var without corresponding ititalization are zero-valued. var int = 0 

:= syntax shorthand for declaring and initializing variables 

*/
